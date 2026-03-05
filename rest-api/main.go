package main

import (
	"fmt"
	"leaning-go-lang/db"
	"leaning-go-lang/routes"
	"log"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("main")
	randomDataGenerate()
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run("localhost:8084")
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
