package main

import (
	db "main/controllers"
	authM "main/middleware"
	"net/http"
)

func HandleGetUUID(w http.ResponseWriter, r *http.Request) {

	database := Client.Database("twclone").Collection("users")
	db.GetUuid(w, r, database)
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	database := Client.Database("twclone").Collection("users")
	db.CreateUser(w, r, database)

}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	database := Client.Database("twclone").Collection("users")
	db.GetUser(w, r, database)

}

func HandleGetTweets(w http.ResponseWriter, r *http.Request) {
	database := Client.Database("twclone").Collection("users")
	limit := r.FormValue("limit")

	authM.AuthMiddleware(w, r, database, func() {
		database = Client.Database("twclone").Collection("tweets")
		db.GetAllTweets(w, r, database, limit)
	})

}

func HandlePostTweet(w http.ResponseWriter, r *http.Request) {
	database := Client.Database("twclone").Collection("users")
	authM.AuthMiddleware(w, r, database, func() {
		database = Client.Database("twclone").Collection("tweets")
		db.AddTweet(w, r, database)
	})
}

func HandleLikeTweet(w http.ResponseWriter, r *http.Request) {
	usersdatabase := Client.Database("twclone").Collection("users")
	authM.AuthMiddleware(w, r, usersdatabase, func() {
		tweetsDatabase := Client.Database("twclone").Collection("tweets")
		db.LikeTweet(w, r, usersdatabase, tweetsDatabase)
	})
}
