package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karokojnr/duka/config"
	"github.com/karokojnr/duka/internal/helper"
	"gorm.io/gorm"
)

type Handler struct {
	App    *fiber.App
	DB     *gorm.DB
	Auth   helper.Auth
	Config config.AppConfig
}
