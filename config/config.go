package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongoDB(uri string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatal(err)
	}

	Client = client
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func CloseMongoDB() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from MongoDB!")
}
