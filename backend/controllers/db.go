package db

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func CreateUser(w http.ResponseWriter, r *http.Request, db *mongo.Collection) {

	// Convert JSON body to Object
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal("decode error")
	}

	// Check if the user already exists
	results := GetFromDB(db, "email", user.Email)

	// If it's a new user
	if results == nil {
		// Give UUID to user
		user.Uid = uuid.NewString()

		res, _ := db.InsertOne(context.TODO(), user)
		fmt.Println(res)

		results := GetFromDB(db, "_id", res.InsertedID)

		userData, _ := json.Marshal(results[0])
		w.Write(userData)
		w.WriteHeader(http.StatusCreated)
	} else {
		// User Exists
		res, _ := json.Marshal(results[0])
		w.Write(res)
		w.WriteHeader(http.StatusAccepted)
	}

}
