package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var app *fiber.App = fiber.New()

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
