package db

import (
	"context"
	"encoding/json"
	"main/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request, db *mongo.Collection) {
	// Convert JSON body to Object
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonRes, _ := json.Marshal(Response{Msg: "Invalid json"})
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(jsonRes)
		return
	}

	// Check if the user already exists
	results := GetFromDB(db, "email", user.Email)

	// If it's a new user
	if results == nil {
		// Give UUID to user
		user.Uid = uuid.NewString()

		res, err := db.InsertOne(context.TODO(), user)
		if err != nil {
			jsonRes, _ := json.Marshal(Response{Msg: "Failed to add tweet to db!"})
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write(jsonRes)
			return
		}

		results := GetFromDB(db, "_id", res.InsertedID)

		userData, _ := json.Marshal(results[0])
		w.WriteHeader(http.StatusCreated)
		w.Write(userData)
	} else {
		// User Exists
		res, _ := json.Marshal(results[0])
		w.WriteHeader(http.StatusAccepted)
		w.Write(res)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, db *mongo.Collection) {
	vars := mux.Vars(r)["id"]
	user := GetFromDB(db, "uid", vars)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		jsonRes, _ := json.Marshal(Response{Msg: "Not found!"})
		w.Write(jsonRes)
	} else {
		jsonUser, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonUser)
	}
}

func GetUuid(w http.ResponseWriter, r *http.Request, db *mongo.Collection) {
	vars := mux.Vars(r)["id"]
	user := GetFromDB(db, "email", vars)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		jsonRes, _ := json.Marshal(Response{Msg: "Not found!"})
		w.Write(jsonRes)
	} else {
		jsonUser, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonUser)
	}
}
