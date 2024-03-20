package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karokojnr/duka/config"
	"github.com/karokojnr/duka/internal/api/rest"
	"github.com/karokojnr/duka/internal/api/rest/handlers"
	"github.com/karokojnr/duka/internal/domain"
	"github.com/karokojnr/duka/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v\n", err)
	}

	log.Println("connected to db🎉")

	// run db migration
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Printf("db migration failed %v\n", err)
	}

	auth := helper.SetupAuth(config.AppSecret)

	handler := &rest.Handler{
		App:  app,
		DB:   db,
		Auth: auth,
	}
	setUpHandlers(handler)

	err = app.Listen(config.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func setUpHandlers(handler *rest.Handler) {
	// User handler
	handlers.SetupUserRoutes(handler)

	// todo: Transactions handler

	// todo: Catalog handler
}
