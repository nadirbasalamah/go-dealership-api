package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-dealership-api/database"
	"github.com/nadirbasalamah/go-dealership-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCar(input model.CarInput, c *fiber.Ctx) model.Car {
	var car model.Car = model.Car{
		Name:  input.Name,
		Brand: input.Brand,
		Year:  input.Year,
		Price: input.Price,
	}

	var collection *mongo.Collection = database.DB.Database.Collection("cars")

	car.ID = ""

	result, err := collection.InsertOne(c.Context(), car)

	if err != nil {
		return model.Car{}
	}

	var filter primitive.D = bson.D{{Key: "_id", Value: result.InsertedID}}
	var createdRecord *mongo.SingleResult = collection.FindOne(c.Context(), filter)

	var createdCar *model.Car = &model.Car{}
	createdRecord.Decode(createdCar)

	return *createdCar
}
