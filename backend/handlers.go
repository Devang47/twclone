package main

import (
	"log"
	db "main/controllers"
	authM "main/middleware"
	"net/http"

	"github.com/gorilla/websocket"
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

	authM.AuthMiddleware(w, r, database, func() {
		database = Client.Database("twclone").Collection("tweets")
		db.GetAllTweets(w, r, database, "all")
	})

}

func HandleGetTweetsOfUser(w http.ResponseWriter, r *http.Request) {
	database := Client.Database("twclone").Collection("users")

	authM.AuthMiddleware(w, r, database, func() {
		database = Client.Database("twclone").Collection("tweets")
		// db.GetAllTweetsByUser(w, r, database)
		db.GetAllTweets(w, r, database, "user")

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

func HandleDeleteTweet(w http.ResponseWriter, r *http.Request) {
	usersdatabase := Client.Database("twclone").Collection("users")
	authM.AuthMiddleware(w, r, usersdatabase, func() {
		tweetsDatabase := Client.Database("twclone").Collection("tweets")
		db.DeleteTweet(w, r, tweetsDatabase)
	})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)

	}

	err = ws.WriteMessage(1, []byte("Hello from twclone!"))
	if err != nil {
		panic(err)
	}

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if string(p) == "tweets-updated" {
			conn.WriteMessage(1, []byte("refresh"))
		}
	}
}
