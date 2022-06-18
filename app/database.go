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

const uri = "mongodb://root:example@localhost:27017/?maxPoolSize=20&w=majority"
const data_path = "../.mongo/data.json"

var client *mongo.Client

var ctx context.Context

var Plan []interface{}

var exercises []bson.M

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

	if err := json.Unmarshal(bytes, &Plan); err != nil {
		panic(err)
	}

	logCollection.InsertMany(*new_ctw, Plan)

	PrintLog("Data inserted")

	file.Close()

	query, err := logCollection.Find(*new_ctw, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = query.All(*new_ctw, &exercises); err != nil {
		panic(err)
	}

	for _, exercise := range exercises {
		workouts = append(workouts, exercise)
	}
}
