package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Plan struct {
	ID        string     `json:"_id"`
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

var Plans []Plan

func ConnectDB() {

	// Must use Mongo 5.0 stable (note: change localhost to 'bibi' when the image is about to upload)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@bibi:27017/"))
	if err != nil {
		panic(err)
	}

	log.Println("Connecting..")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	PrintLog("Connected to MongoDB")
	setupMongo(client, &ctx)

}

func setupMongo(cl *mongo.Client, new_ctw *context.Context) {

	logCollection := *cl.Database("database").Collection("temp")

	PrintLog("Connecting to Database")

	file, err := os.Open("../data.json")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(bytes, &Plans); err != nil {
		panic(err)
	}

	for _, item := range Plans {
		logCollection.InsertOne(*new_ctw, item)
	}

	PrintLog("Data inserted")

	file.Close()
}
