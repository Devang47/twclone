package db

import (
	"backend/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTweet(w http.ResponseWriter, r *http.Request, db *mongo.Collection) {

	// Convert JSON body to Object
	var tweet models.Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		jsonRes, _ := json.Marshal(Response{Msg: "Invalid json"})
		w.Write(jsonRes)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	currentTime := time.Now()
	tweet.PublishedOn = currentTime.Format("2006-01-02 15:04:05.000000")

	_, err := db.InsertOne(context.TODO(), tweet)
	if err != nil {
		jsonRes, _ := json.Marshal(Response{Msg: "Failed to add tweet to db!"})
		w.Write(jsonRes)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	jsonRes, _ := json.Marshal(Response{Msg: "Done"})
	w.Write(jsonRes)
	w.WriteHeader(http.StatusOK)

}

func GetAllTweets(w http.ResponseWriter, r *http.Request, db *mongo.Collection) {

	cursor, err := db.Find(context.TODO(), bson.D{primitive.E{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	jsonTweets, err := json.Marshal(results)
	if err != nil {
		jsonRes, _ := json.Marshal(Response{Msg: "Invalid json!"})
		w.Write(jsonRes)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	w.Write(jsonTweets)
	w.WriteHeader(http.StatusOK)
}
