package main

import (
	"damo-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	log.Println("getEvents")
	server.Run("localhost:8081")
}

func createEvent(context *gin.Context) {
	context.JSON(http.StatusOK, model.Event{})
	var event model.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event"})
		return
	}
	event.ID = "1"
	event.UserId = "1"
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "status": "success", "code": 200})
	log.Println("createEvent")
}

func getEvents(c *gin.Context) {
	log.Println("getEvents")
	c.JSON(http.StatusOK, model.EventsList())
}
