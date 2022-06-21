package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://root:example@localhost:27017/?maxPoolSize=20&w=majority"
const data_path = "../.mongo/data.json"

var client *mongo.Client

var ctx context.Context

var Plans []Plan

type Plan struct {
	ID        string     `json:"_id"`
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

func ConnectDB() {

	// Must use Mongo 5.0 stable (note: change localhost to 'bibi' when the image is about to upload)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	PrintLog("Connected to MongoDB")

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	setupMongo(client, &ctx)

}

func setupMongo(cl *mongo.Client, new_ctw *context.Context) {

	logCollection := *cl.Database("database").Collection("temp")

	PrintLog("Connecting to Database")

	file, err := os.Open(data_path)
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
