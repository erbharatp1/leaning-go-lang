package main

import (
	"fmt"
	"leaning-go-lang/db"
	"leaning-go-lang/model"
	"log"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("main")
	randomDataGenerate()
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:name", getByName)

	server.Run("localhost:8082")
}

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

func randomDataGenerate() {
	log.Println("random Data Generator")
	fmt.Println(randomdata.FullName(randomdata.RandomGender))
	fmt.Println(randomdata.Email())
	fmt.Println(randomdata.Address())
	fmt.Println(randomdata.City())
	fmt.Println(randomdata.State(randomdata.Large))
	fmt.Println(randomdata.Country(randomdata.FullCountry))
	fmt.Println(randomdata.PostalCode("SE"))
	fmt.Println(randomdata.PhoneNumber())
	fmt.Println(randomdata.IpV4Address())
	fmt.Println(randomdata.Day())
	fmt.Println(randomdata.Month())

}
