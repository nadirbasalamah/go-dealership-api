package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-dealership-api/model"
	"github.com/nadirbasalamah/go-dealership-api/service"
)

func CreateCar(c *fiber.Ctx) error {
	var carInput *model.CarInput = new(model.CarInput)

	if err := c.BodyParser(carInput); err != nil {
		c.Status(http.StatusUnprocessableEntity).SendString(err.Error())
		return err
	}

	var errors []*model.ErrorResponse = carInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(errors)
	}

	var createdCar model.Car = service.CreateCar(*carInput, c)

	return c.Status(http.StatusCreated).JSON(createdCar)
}
