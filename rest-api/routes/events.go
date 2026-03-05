package routes

import (
	"leaning-go-lang/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getByName(context *gin.Context) {
	name := context.Param("name")
	log.Println("getByName")
	event, err := model.FindByName(name)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	context.JSON(http.StatusOK, event)
	log.Println(event)
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

func updateEventByID(context *gin.Context) {
	idStr := context.Param("id")
	log.Printf("updateEventByID: attempting to update event with ID: %s", idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("updateEventByID: invalid ID format: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID format", "id": idStr})
		return
	}

	event, err := model.FindByID(id)
	if err != nil {
		log.Printf("updateEventByID: event not found: %v", err)
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found", "id": id})
		return
	}

	var updatedEvent model.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		log.Printf("updateEventByID: failed to bind JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event data. Please ensure all required fields (name, location, description) are provided in the correct format.", "details": err.Error()})
		return
	}

	updatedEvent.ID = event.ID
	updatedEvent.UserId = event.UserId
	if err := updatedEvent.Update(); err != nil {
		log.Printf("updateEventByID: failed to update event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "status": "success", "code": 200, "event": updatedEvent})
	log.Printf("updateEventByID: successfully updated event ID %d", updatedEvent.ID)
}
