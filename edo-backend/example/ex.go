package example

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Config structure for database configuration
type Config struct {
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
}

func LoadConfig() (Config, error) {
	var config Config
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}

func ConnectDB(config Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DBUser, config.DBPassword, config.DBName)
	return sql.Open("postgres", connStr)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(db *sql.DB, username, password string) error {
	passwordHash, err := hashPassword(password)
	if err != nil {
		return err
	}
	query, args, err := squirrel.
		Insert("private.users").
		Columns("username", "password_hash").
		Values(username, passwordHash).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	_, err = db.Exec(query, args...)

	return err
}

func authenticateUser(db *sql.DB, username, password string) (bool, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT password_hash FROM private.users WHERE username = $1", username).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}
	return checkPasswordHash(password, hashedPassword), nil
}

// Struct to handle login requests
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Bad request",
			})
		}

		authenticated, err := authenticateUser(db, req.Username, req.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		if authenticated {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Authentication successful!",
			})
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authentication failed!",
			})
		}
	}
}
