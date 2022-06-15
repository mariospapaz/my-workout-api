package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
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

func ConnectDB() *mongo.Collection {
	// Must use Mongo 5.0 stable
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@bibi:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")
	logCollection := client.Database("my-workout").Collection("temp")
	return logCollection
}

func main() {

	log.Println("##############################################")
	log.Println("###       Connecting to Database        ######")
	log.Println("##############################################")

	col := ConnectDB()

	log.Println(col)

	router := gin.Default()
	RouterPaths(router)
}
