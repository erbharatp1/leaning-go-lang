package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running correctly"})
	})
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:name", getByName)
	server.PUT("/events/id/:id", updateEventByID)
	server.POST("/signup", signup)
	server.POST("/login", login)

}
