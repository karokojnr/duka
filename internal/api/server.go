package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karokojnr/duka/config"
	"github.com/karokojnr/duka/internal/api/rest"
	"github.com/karokojnr/duka/internal/api/rest/handlers"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	handler := &rest.RestHandler{
		App: app,
	}
	setUpHandlers(handler)

	err := app.Listen(config.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func setUpHandlers(handler *rest.RestHandler) {
	// User handler
	handlers.SetupUserRoutes(handler)

	// todo: Transactions handler

	// todo: Catalog handler
}
