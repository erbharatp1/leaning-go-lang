package main

import (
	"damo-go/db"
	"damo-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("main")
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	log.Println("getEvents")
	server.Run("localhost:8082")
}

func createEvent(context *gin.Context) {
	var event model.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event"})
		return
	}
	event.UserId = "1"
	if err := event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "status": "success", "code": 200, "event": event})
	log.Println("createEvent")
}

func getEvents(c *gin.Context) {
	log.Println("getEvents")
	c.JSON(http.StatusOK, model.EventsList())
}
