package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5/middleware"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"

	db "backend/controllers"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	mux.CORSMethodMiddleware(r)

	client := db.ConnectToDB()
	db := client.Database("twclone").Collection("temp")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars)

		doc := bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}, {"id", vars["id"]}}

		db.InsertOne(context.TODO(), doc)

		w.Write([]byte("Hello world "))
	}).Methods("GET").Schemes("http")

	d := color.New(color.FgGreen, color.Bold)
	c := color.New(color.FgGreen)
	c.Printf("Started backend server at: ")
	d.Println("5000")
	d.Println("	➜  http://localhost:5000/")

	http.Handle("/", r)
	s := http.Server{Addr: ":5000", Handler: r}
	s.ListenAndServe()
}
