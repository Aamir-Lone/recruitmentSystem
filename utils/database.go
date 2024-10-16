package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectDB connects to the MongoDB database
func ConnectDB() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb+srv://Aamirlone:Aamirlone@cluster0.1zwiz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}

// DisconnectDB disconnects from the MongoDB database
func DisconnectDB() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB!")
}

/*package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDatabase() {
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("YOUR_MONGODB_URI"))
	if err != nil {
		log.Fatal(err)
	}
}
*/
