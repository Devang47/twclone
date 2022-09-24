package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"

	db "main/controllers"
)

var Client *mongo.Client

func main() {
	godotenv.Load()

	Client = db.ConnectToDB()
	r := SetupRoutes()

	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

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
