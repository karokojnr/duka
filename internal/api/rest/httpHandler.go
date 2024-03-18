package rest

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	App *fiber.App
	DB  *gorm.DB
}
