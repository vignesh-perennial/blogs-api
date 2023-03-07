package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"blogs_api/utils/registrybuilder"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var mongoClient *mongo.Client

func Db() *mongo.Client {
	once.Do(func() {

		uri := "mongodb://localhost:27017"

		//creating context
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(uri).SetRegistry(registrybuilder.MongoRegistry)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			fmt.Println("Error connecting to mongodb: ", err)
		}

		mongoClient = client
		fmt.Println("Successfully established connection to mongodb!!")

	})
	return mongoClient
}
