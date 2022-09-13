package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	db "backend/controllers"
	authM "backend/middleware"
)

func main() {
	godotenv.Load()

	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.Use(handlers.RecoveryHandler())
	r.Use(mux.CORSMethodMiddleware(r))

	client := db.ConnectToDB()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r.HandleFunc("/get-uuid/{id}", func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("twclone").Collection("users")
		db.GetUuid(w, r, database)

	}).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("twclone").Collection("users")
		db.CreateUser(w, r, database)

	}).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("twclone").Collection("users")
		db.GetUser(w, r, database)

	}).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/get-tweets", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Headers", " Authorization")
		database := client.Database("twclone").Collection("users")
		authM.AuthMiddleware(w, r, database, func() {
			database = client.Database("twclone").Collection("tweets")
			db.GetAllTweets(w, r, database)
		})

	}).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	r.HandleFunc("/post-tweet", func(w http.ResponseWriter, r *http.Request) {
		database := client.Database("twclone").Collection("users")
		authM.AuthMiddleware(w, r, database, func() {
			database = client.Database("twclone").Collection("tweets")
			db.AddTweet(w, r, database)
		})

	}).Methods("POST", http.MethodOptions)

	fmt.Printf("Started backend server at: ")
	fmt.Println("5000")
	fmt.Println("➜  http://localhost:5000/")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Content-type", "Authorization"},
	})

	handler := c.Handler(r)
	http.ListenAndServe(":5000", handler)
	http.Handle("/", r)
}
