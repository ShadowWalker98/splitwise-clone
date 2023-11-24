package databasehandler

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"splitwise-clone/types"
)

func Connect() *mongo.Client {
	// Create a new client and connect to the server
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://admin:" + os.Getenv("password") + "@cluster0.xj356kx.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return client
}

func InsertItem(client *mongo.Client, item types.ExpenseItem) {
	result, err := client.Database("splitwise-clone").Collection("expenses").InsertOne(context.TODO(), item)

	if err != nil {
		panic(err)
	}

	fmt.Println("Added an item. Result: {}", result.InsertedID)
}
