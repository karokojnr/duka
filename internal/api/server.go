package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karokojnr/duka/config"
	"log"
	"net/http"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", healthCheck)

	err := app.Listen(config.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Server running",
	})
}
