package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllWorkouts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Plans)
}

func GetDayWorkout(c *gin.Context) {
	id := c.Param("id")
	n_id, _ := strconv.Atoi(id)
	log.Println(n_id)

	if n_id > 7 || n_id < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "each week has 7 days."})
		return
	}

	if Plans[n_id-1].Exercises == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "break"})
		return
	}

	c.IndentedJSON(http.StatusOK, Plans[n_id-1])
}

func PrintLog(line string) {
	log.Println("##############################################")
	log.Printf("###       %s                           ######\n", line)
	log.Println("##############################################")
}

func main() {
	ConnectDB()

	router := gin.Default()

	router.GET("/api/workout", GetAllWorkouts)
	router.GET("/api/workout/:id", GetDayWorkout)
	router.Run("localhost:25585")
}
