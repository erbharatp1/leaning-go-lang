package main

import (
	"leaning-go-lang/db"
	"leaning-go-lang/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("main")
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run("localhost:8084")
}
