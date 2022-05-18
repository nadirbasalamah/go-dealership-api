package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-dealership-api/database"
)

func main() {
	var app *fiber.App = fiber.New()

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
