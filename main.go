package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/KayacChang/API_Server/user"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func main() {
	r := httprouter.New()

	c := initDB()

	db := c.Database("test")

	user.Mount(r, db)

	http.ListenAndServe(":8080", r)
}
