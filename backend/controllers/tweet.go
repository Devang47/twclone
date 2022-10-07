package db

import (
	"context"
	"encoding/json"
	"log"
	"main/models"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(jsonRes)
		return
	}

	// currentTime := time.Now()
	// tweet.PublishedOn = currentTime.Format("2006-01-02 15:04:05.000000")
	tweet.Id = uuid.NewString()

	_, err := db.InsertOne(context.TODO(), tweet)
	if err != nil {
		jsonRes, _ := json.Marshal(Response{Msg: "Failed to add tweet to db!"})
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write(jsonRes)
		return
	}

	jsonRes, _ := json.Marshal(Response{Msg: "Done"})
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)

}

func GetAllTweets(w http.ResponseWriter, r *http.Request, db *mongo.Collection, method string) {

	var cursor *mongo.Cursor
	var err error
	if method == "all" {
		limit := r.FormValue("limit")

		i, err := strconv.Atoi(limit)
		if err != nil {
			panic(err)
		}

		opts := options.Find().SetSort(bson.D{{Key: "published_on", Value: -1}}).SetLimit(int64(i))

		cursor, err = db.Find(context.TODO(), bson.D{primitive.E{}}, opts)
		if err != nil {
			log.Fatal(err)
		}

	} else if method == "user" {
		vars := mux.Vars(r)["id"]

		opts := options.Find().SetSort(bson.D{{Key: "published_on", Value: -1}})

		cursor, err = db.Find(context.TODO(), bson.D{primitive.E{Key: "author_uid", Value: vars}}, opts)
		if err != nil {
			log.Fatal(err)
		}
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	jsonTweets, err := json.Marshal(results)
	if err != nil {
		jsonRes, _ := json.Marshal(Response{Msg: "Invalid json!"})
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(jsonRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonTweets)
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
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
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
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func DeleteTweet(w http.ResponseWriter, r *http.Request, db *mongo.Collection) {

	var tweet models.Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		jsonRes, _ := json.Marshal(Response{Msg: "Invalid json"})
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(jsonRes)
		return
	}

	UID := r.Header.Get("Authorization")
	if tweet.AuthorUid != UID {
		jsonRes, _ := json.Marshal(Response{Msg: "Authorization failed!"})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonRes)
		return

	}

	_, err := db.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "id", Value: tweet.Id}})
	if err != nil {
		jsonRes, _ := json.Marshal(Response{Msg: "Failed to add tweet to db!"})
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write(jsonRes)
		return
	}

	jsonRes, _ := json.Marshal(Response{Msg: "Done"})
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)

}
