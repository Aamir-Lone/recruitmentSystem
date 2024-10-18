package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the MongoDB URI from the .env file
	mongoURI := os.Getenv("MONGODB_URI")
	// Create a new MongoDB client using the updated method

	clientOptions := options.Client().ApplyURI(mongoURI)

	// Use mongo.Connect instead of mongo.NewClient
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection with a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	Client = client
	fmt.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("recruitment_db").Collection(collectionName)
}
