package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-dealership-api/controller"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/cars", controller.GetAllCars)
	app.Get("api/cars/:id", controller.GetCar)
	app.Post("/api/cars", controller.CreateCar)
	app.Put("/api/cars/:id", controller.UpdateCar)
}
