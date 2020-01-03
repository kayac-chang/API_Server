package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Serve(r *httprouter.Router, db *mongo.Database) *httprouter.Router {

	r.POST("/users", create(db))

	return r
}

func create(db *mongo.Database) httprouter.Handle {

	tb := db.Collection("users")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		//	Create User
		u := user{
			ID: primitive.NewObjectID(),
		}

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			//TODO
		}

		//	Add to Database
		ctx := context.TODO()

		_, err = tb.InsertOne(ctx, u)

		if err != nil {
			//TODO
		}

		//	Response to Client
		w.WriteHeader(http.StatusCreated)

		sendJSON(w, u)
	}
}
