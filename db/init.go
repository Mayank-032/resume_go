package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ResumeDataCollection *mongo.Collection
func ConnectDB(connectionURI string) {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	ResumeDataCollection = client.Database("resumeDB").Collection("resume")
}