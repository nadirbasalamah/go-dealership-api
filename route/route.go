package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-dealership-api/controller"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/cars", controller.CreateCar)
}
