package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-dealership-api/database"
	"github.com/nadirbasalamah/go-dealership-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllCars(c *fiber.Ctx) []model.Car {
	var query primitive.D = bson.D{{}}

	cursor, err := database.DB.Database.Collection("cars").Find(c.Context(), query)
	if err != nil {
		return []model.Car{}
	}

	var cars []model.Car = make([]model.Car, 0)

	if err := cursor.All(c.Context(), &cars); err != nil {
		return []model.Car{}
	}

	return cars
}

func GetCar(id string, c *fiber.Ctx) model.Car {
	carID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Car{}
	}

	var query primitive.D = bson.D{{Key: "_id", Value: carID}}
	var collection *mongo.Collection = database.DB.Database.Collection("cars")

	var carData *mongo.SingleResult = collection.FindOne(c.Context(), query)

	var car *model.Car = &model.Car{}
	carData.Decode(car)

	return *car
}

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
