package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetAllWorkouts(ctx *gin.Context) {

}

func GetDayWorkout(ctx *gin.Context) {

}

func RouterPaths(engine *gin.Engine) {
	engine.GET("/api/workout", GetAllWorkouts)
	engine.GET("/api/workout/:day", GetDayWorkout)
	engine.Run("localhost:8080")
}

func PrintLog(line string) {
	log.Println("##############################################")
	log.Printf("###       %s                           ######\n", line)
	log.Println("##############################################")
}

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

	databases, err := client.ListDatabases(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	file, err := os.Open(data_path)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)

	var Plan []interface{}

	if err := json.Unmarshal(bytes, &Plan); err != nil {
		panic(err)
	}

	log.Println(databases)

	logCollection.InsertMany(ctx, Plan)

	PrintLog("Data inserted")

	file.Close()
}

func main() {

	ConnectDB()

	//router := gin.Default()
	//RouterPaths(router)
}
