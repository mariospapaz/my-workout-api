package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() {

	// Must use Mongo 5.0 stable (note: change localhost to 'bibi' when the image is about to upload)

	PrintLog("Connecting to Database")

	const uri = "mongodb://root:example@localhost:27017/?maxPoolSize=20&w=majority"
	const data_path = "../.mongo/data.json"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

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

	logCollection := client.Database("database").Collection("temp")

	file, err := os.Open(data_path)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)

	var Plan []interface{}

	if err := json.Unmarshal(bytes, &Plan); err != nil {
		panic(err)
	}

	logCollection.InsertMany(ctx, Plan)

	PrintLog("Data inserted")

	file.Close()

	query, err := logCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	var exercises []bson.M
	if err = query.All(ctx, &exercises); err != nil {
		panic(err)
	}

	for _, exercise := range exercises {
		workouts = append(workouts, exercise)
	}
}
