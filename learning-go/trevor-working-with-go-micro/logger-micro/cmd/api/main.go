package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = 80
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo-container:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
}

func main() {
	// connect to mongo db
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	// create a context for disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()
}

func connectToMongo() (*mongo.Client, error) {
	// connect 	connection options
	clientOptions := options.Client()

	clientOptions.SetAuth(options.Credential{
		Username: "root",
		Password: "root",
	})

	c, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println("Error connecting to mongo: ", err)
		return nil, err
	}

	return c, nil
}
