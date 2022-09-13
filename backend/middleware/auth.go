package middleware

import (
	db "backend/controllers"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func AuthMiddleware(w http.ResponseWriter, r *http.Request, dbI *mongo.Collection, next func()) {
	UID := r.Header.Get("Authorization")

	if UID == "" {
		jsonRes, _ := json.Marshal(db.Response{Msg: "Authorization header not found"})
		w.WriteHeader(http.StatusForbidden)
		w.Write(jsonRes)
		panic("Authorization header not found")

	}

	user := db.GetFromDB(dbI, "uid", UID)

	if user == nil {
		jsonRes, _ := json.Marshal(db.Response{Msg: "Invalid authorization header"})
		w.WriteHeader(http.StatusForbidden)
		w.Write(jsonRes)
		panic("Invalid authorization header")
	}

	next()
}
