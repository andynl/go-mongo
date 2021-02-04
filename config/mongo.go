package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDB() (*mongo.Database, error) {
	// host := os.Getenv("MONGO_HOST")
	// dbName := os.Getenv("MONGO_DB_NAME")
	host := "mongodb://localhost:27017"
	dbName := "tutorial1"

	ctx, cancel := NewMongoContext()
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
