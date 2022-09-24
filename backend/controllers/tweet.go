package db

import (
	"context"
	"encoding/json"
	"log"
	"main/models"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	tweet.Id = uuid.NewString()

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

func GetAllTweets(w http.ResponseWriter, r *http.Request, db *mongo.Collection, limit string) {

	i, err := strconv.Atoi(limit)
	if err != nil {
		panic(err)
	}

	opts := options.Find().SetSort(bson.D{{Key: "published_on", Value: -1}}).SetLimit(int64(i))

	cursor, err := db.Find(context.TODO(), bson.D{primitive.E{}}, opts)
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

func LikeTweet(w http.ResponseWriter, r *http.Request, usersDatabase *mongo.Collection, tweetsDatabase *mongo.Collection) {
	tweetId := r.FormValue("id")
	UID := r.Header.Get("Authorization")
	user := GetFromDB(usersDatabase, "uid", UID)

	tweet := GetTweetFromDB(tweetsDatabase, "id", tweetId)

	for index, i := range tweet[0].Likes.LikedBy {
		if user[0]["username"].(string) == i {
			tweetsDatabase.UpdateOne(context.TODO(), bson.M{"id": tweetId}, bson.D{
				{Key: "$set", Value: bson.D{{Key: "likes", Value: bson.D{
					{
						Key:   "total_likes",
						Value: tweet[0].Likes.TotalLikes - 1,
					},
					{
						Key:   "liked_by",
						Value: remove(tweet[0].Likes.LikedBy, index),
					},
				}}}},
			})

			jsonRes, _ := json.Marshal(Response{Msg: "Done"})
			w.Write(jsonRes)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	tweetsDatabase.UpdateOne(context.TODO(), bson.M{"id": tweetId}, bson.D{
		{Key: "$set", Value: bson.D{{Key: "likes", Value: bson.D{
			{
				Key:   "total_likes",
				Value: tweet[0].Likes.TotalLikes + 1,
			},
			{
				Key:   "liked_by",
				Value: append(tweet[0].Likes.LikedBy, user[0]["username"].(string)),
			},
		}}}},
	})

	jsonRes, _ := json.Marshal(Response{Msg: "Done"})
	w.Write(jsonRes)
	w.WriteHeader(http.StatusOK)
	return
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
