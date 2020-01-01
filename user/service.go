package user

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

	insert := func(u user) (*mongo.InsertOneResult, error) {
		ctx := context.TODO()

		res, err := tb.InsertOne(ctx, u)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		//	1. Create User
		u := user{
			ID: primitive.NewObjectID(),
		}

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			//TODO: Parse error
		}

		//	2. Add to Database
		_, err = insert(u)

		if err != nil {
			//TODO: Insert error
		}

		//	3. Response to Client
		w.WriteHeader(http.StatusCreated)

		sendJSON(w, u)
	}
}
