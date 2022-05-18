package database

import (
	"context"
	"time"

	"github.com/nadirbasalamah/go-dealership-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var DB MongoInstance

func Connect() error {
	client, _ := mongo.NewClient(options.Client().ApplyURI(config.Config("MONGO_URI")))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := client.Connect(ctx)
	var db *mongo.Database = client.Database(config.Config("DATABASE_NAME"))

	if err != nil {
		return err
	}

	DB = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}
