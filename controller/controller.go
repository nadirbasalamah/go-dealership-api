package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-dealership-api/model"
	"github.com/nadirbasalamah/go-dealership-api/service"
)

func GetAllCars(c *fiber.Ctx) error {
	var cars []model.Car = service.GetAllCars(c)

	return c.JSON(cars)
}

func GetCar(c *fiber.Ctx) error {
	var carID string = c.Params("id")

	var car model.Car = service.GetCar(carID, c)
	if car.ID == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Car not found",
		})
	}

	return c.JSON(car)
}

func CreateCar(c *fiber.Ctx) error {
	var carInput *model.CarInput = new(model.CarInput)

	if err := c.BodyParser(carInput); err != nil {
		return c.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}

	var errors []*model.ErrorResponse = carInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}

	var createdCar model.Car = service.CreateCar(*carInput, c)

	return c.Status(http.StatusCreated).JSON(createdCar)
}

func UpdateCar(c *fiber.Ctx) error {
	var carID string = c.Params("id")

	var carInput *model.CarInput = new(model.CarInput)

	if err := c.BodyParser(carInput); err != nil {
		return c.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}

	var errors []*model.ErrorResponse = carInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}

	var updatedCar model.Car = service.UpdateCar(carID, *carInput, c)

	return c.JSON(updatedCar)
}

func DeleteCar(c *fiber.Ctx) error {
	var carID string = c.Params("id")

	var result bool = service.DeleteCar(carID, c)
	if result == false {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Car not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Car deleted",
	})
}
