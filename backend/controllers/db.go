package db

import (
	"context"
	"fmt"
	"log"
	"main/models"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Msg string `json:"msg"`
}

func ConnectToDB() *mongo.Client {
	URI := os.Getenv("URI")

	if err := godotenv.Load(); err != nil {
		fmt.Println("No ENV file was found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}

	return client
}

func GetFromDB(db *mongo.Collection, key string, value any) []bson.M {

	cursor, err := db.Find(context.TODO(), bson.D{primitive.E{Key: key, Value: value}})
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}

func GetTweetFromDB(db *mongo.Collection, key string, value any) []models.Tweet {

	cursor, err := db.Find(context.TODO(), bson.D{primitive.E{Key: key, Value: value}})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Tweet
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}
