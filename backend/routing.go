package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.Use(handlers.RecoveryHandler())
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/api/get-uuid/{id}", HandleGetUUID).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/api/create-user", HandleCreateUser).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/user/{id}", HandleGetUser).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/api/get-tweets", HandleGetTweets).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/api/get-tweets/{id}", HandleGetTweetsOfUser).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/api/post-tweet", HandlePostTweet).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/like-tweet", HandleLikeTweet).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/api/delete-tweet", HandleDeleteTweet).Methods(http.MethodPost, http.MethodOptions)

	return r
}
