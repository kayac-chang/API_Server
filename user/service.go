package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	ID     primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
}

func Mount(r *httprouter.Router, db *mongo.Database) *httprouter.Router {

	r.GET("/user/:id", get(db))

	r.POST("/user", create(db))

	// r.DELETE("/user/:id", delete(db))

	return r
}

func get(db *mongo.Database) httprouter.Handle {

	collection := db.Collection("users")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		u := user{}

		idStr := p.ByName("id")

		docID, err := primitive.ObjectIDFromHex(idStr)

		if err != nil {
			log.Fatal(err)
		}

		filter := bson.M{"_id": docID}

		err = collection.FindOne(context.TODO(), filter).Decode(&u)

		if err != nil {
			log.Fatal(err)
		}

		uj, err := json.Marshal(u)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s\n", uj)
	}
}

func create(db *mongo.Database) httprouter.Handle {

	collection := db.Collection("users")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		u := user{
			ID: primitive.NewObjectID(),
		}

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			log.Fatal(err)
		}

		_, err = collection.InsertOne(context.TODO(), u)

		if err != nil {
			log.Fatal(err)
		}

		uj, err := json.Marshal(u)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s\n", uj)
	}
}

func delete(db *mongo.Database) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Write code to delete user\n")
	}
}
