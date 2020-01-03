package mongo

import (
	"github.com/KayacChang/API_Server/core"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	host = "localhost"
	port = 27017
)

var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	_client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = _client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Successed...")

	client = _client
}

func New(name string) *mongo.Database {
	return client.Database(name)
}
