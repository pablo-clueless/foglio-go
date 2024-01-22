package common

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var helpers Helpers

func ConnectMongo() mongo.Client {
	uri := "mongodb+srv://smsnmicheal:1hIfbQtBHx1rkbwV@cluster0.6zpab4p.mongodb.net/foglio?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		helpers.ErrorLogger.Panicln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		helpers.ErrorLogger.Fatalln(err)
	}
	fmt.Println("Connected to MongoDB!")

	return *client
}

var DB mongo.Client = ConnectMongo()

func GetCollection(client mongo.Client, collection string) *mongo.Collection {
	return client.Database("foglio").Collection(collection)
}
