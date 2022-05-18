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

	cursor, err := database.GetCollection("cars").Find(c.Context(), query)
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
	var collection *mongo.Collection = database.GetCollection("cars")

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

	var collection *mongo.Collection = database.GetCollection("cars")

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

func UpdateCar(id string, input model.CarInput, c *fiber.Ctx) model.Car {
	carID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Car{}
	}

	var query primitive.D = bson.D{{Key: "_id", Value: carID}}
	var update primitive.D = bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "brand", Value: input.Brand},
			{Key: "name", Value: input.Name},
			{Key: "year", Value: input.Year},
			{Key: "price", Value: input.Price},
		},
	}}

	var collection *mongo.Collection = database.GetCollection("cars")

	var updateResult *mongo.SingleResult = collection.FindOneAndUpdate(c.Context(), query, update)

	if updateResult.Err() != nil {
		if err == mongo.ErrNoDocuments {
			return model.Car{}
		}
		return model.Car{}
	}

	var car model.Car = GetCar(id, c)
	return car
}

func DeleteCar(id string, c *fiber.Ctx) bool {
	carID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}

	var query primitive.D = bson.D{{Key: "_id", Value: carID}}
	var collection *mongo.Collection = database.GetCollection("cars")

	result, err := collection.DeleteOne(c.Context(), query)
	var isFailed bool = err != nil || result.DeletedCount < 1

	if isFailed {
		return !isFailed
	}

	return true
}
