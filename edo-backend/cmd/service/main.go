package main

import (
	ex "admin-panel/example"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config, err := ex.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := ex.ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Register a new user for demonstration
	err = ex.RegisterUser(db, "exampleUser", "examplePassword")
	if err != nil {
		log.Fatalf("Error registering user: %v", err)
	}

	// Set up the login route
	app.Post("/login", ex.LoginHandler(db))

	fmt.Println("Starting server on :8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
