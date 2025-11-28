package main

import (
	adapter "admin-panel/internal/adapter"
	"admin-panel/internal/api/handlers/get_user_by_id"
	"admin-panel/internal/domain/user/repository"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	db := adapter.NewProvider().ConnectDB()
	defer db.Close()

	userStorage := repository.NewUserStorage(db)
	userHandler := get_user_by_id.NewHandler(userStorage)
	app := fiber.New()
	app.Get("/users", userHandler.Handle)

	log.Println("Сервер запущен на :8080")
	log.Fatal(app.Listen(":8080"))
}
