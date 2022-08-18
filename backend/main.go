package main

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {

	d := color.New(color.FgGreen, color.Bold)
	c := color.New(color.FgGreen)
	c.Printf("Started backend server at: ")
	d.Println("8080")

	d.Println("	➜  http://localhost:8080/")

	// color.Bold

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(cors.Default().Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	s := http.Server{Addr: ":8080", Handler: r}
	s.ListenAndServe()
}
