package db

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Client {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No ENV file was found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		panic(err)
	}

	return client
}
