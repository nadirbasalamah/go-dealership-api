package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-dealership-api/model"
	"github.com/nadirbasalamah/go-dealership-api/service"
)

func GetAllCars(c *fiber.Ctx) error {
	var cars []model.Car = service.GetAllCars(c)

	var response model.Response[[]model.Car] = model.Response[[]model.Car]{
		Success: true,
		Message: "All cars data",
		Data:    cars,
	}

	return c.JSON(response)
}

func GetCar(c *fiber.Ctx) error {
	var carID string = c.Params("id")

	var car model.Car = service.GetCar(carID, c)
	if car.ID == "" {
		return c.Status(http.StatusNotFound).JSON(model.Response[any]{
			Success: false,
			Message: "Car data not found",
		})
	}

	return c.JSON(model.Response[model.Car]{
		Success: true,
		Message: "Car data found",
		Data:    car,
	})
}

func CreateCar(c *fiber.Ctx) error {
	var carInput *model.CarInput = new(model.CarInput)

	if err := c.BodyParser(carInput); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(model.Response[any]{
			Success: false,
			Message: "Please fill all the required fields",
		})
	}

	var errors []*model.ErrorResponse = carInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}

	var createdCar model.Car = service.CreateCar(*carInput, c)

	return c.Status(http.StatusCreated).JSON(model.Response[model.Car]{
		Success: true,
		Message: "Car data added",
		Data:    createdCar,
	})
}

func UpdateCar(c *fiber.Ctx) error {
	var carID string = c.Params("id")

	var carInput *model.CarInput = new(model.CarInput)

	if err := c.BodyParser(carInput); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(model.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	var errors []*model.ErrorResponse = carInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}

	var updatedCar model.Car = service.UpdateCar(carID, *carInput, c)

	return c.JSON(model.Response[model.Car]{
		Success: true,
		Message: "Car data updated",
		Data:    updatedCar,
	})
}

func DeleteCar(c *fiber.Ctx) error {
	var carID string = c.Params("id")

	var isSuccess bool = service.DeleteCar(carID, c)
	if !isSuccess {
		return c.Status(http.StatusNotFound).JSON(model.Response[any]{
			Success: false,
			Message: "Car data not found",
		})
	}

	return c.JSON(model.Response[any]{
		Success: true,
		Message: "Car data deleted",
	})
}
