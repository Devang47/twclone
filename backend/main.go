package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	db "backend/controllers"
)

func main() {
	godotenv.Load()

	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.Use(handlers.RecoveryHandler())
	// r.Use(mux.CORSMethodMiddleware(r))

	client := db.ConnectToDB()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	/**
		{
	  "given_name": "Devang",
	  "family_name": "Saklani",
	  "nickname": "devangsaklani",
	  "name": "Devang Saklani",
	  "picture": "https://lh3.googleusercontent.com/a-/AFdZucr7NVUE9C66t4IYhwq-oIkyrDbYDkR-h6DhhWB5lwc=s96-c",
	  "locale": "en",
	  "updated_at": "2022-08-28T15:31:39.108Z",
	  "email": "devangsaklani@gmail.com",
	  "email_verified": true,
	  "sub": "google-oauth2|101176798958313563407"
	}
	*/

	r.HandleFunc("/get-uuid/{id}", func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("twclone").Collection("users")
		db.GetUuid(w, r, database)
	}).Methods("GET")

	r.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("twclone").Collection("users")
		db.CreateUser(w, r, database)
	}).Methods("POST")

	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("twclone").Collection("users")
		db.GetUser(w, r, database)

	}).Methods("GET")

	r.HandleFunc("/get-tweets", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		w.Header().Set("access-control-allow-origin", "https://twclone.saklani.dev")
		w.Header().Set("access-control-allow-headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET")

		database := client.Database("twclone").Collection("tweets")
		db.GetAllTweets(w, r, database)

	}).Methods("GET", http.MethodOptions)

	r.HandleFunc("/post-tweet", func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("twclone").Collection("tweets")
		db.AddTweet(w, r, database)
	}).Methods("POST")

	fmt.Printf("Started backend server at: ")
	fmt.Println("5000")
	fmt.Println("➜  http://localhost:5000/")

	http.ListenAndServe(":5000", handlers.CORS()(r))
}
