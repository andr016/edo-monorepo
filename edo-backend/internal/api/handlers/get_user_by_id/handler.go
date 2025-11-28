package get_user_by_id

import (
	storage "admin-panel/internal/domain/user/repository"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Handler struct {
	userStorage *storage.UserStorage
}

func NewHandler(us *storage.UserStorage) *Handler {
	return &Handler{userStorage: us}
}

func (h *Handler) Handle(c *fiber.Ctx) error {
	users, err := h.userStorage.GetAll()
	if err != nil {
		log.Printf("Ошибка в GetUsers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "не удалось получить пользователей",
		})
	}
	return c.JSON(users)
}
