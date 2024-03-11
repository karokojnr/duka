package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Welcome to Duka")

	app := fiber.New()

	// routes

	app.Listen("localhost:4000")
}
