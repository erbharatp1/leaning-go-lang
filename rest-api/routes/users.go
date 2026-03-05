package routes

import (
	"leaning-go-lang/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse user"})
		return
	}
	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}
	// For security reasons, don't return the hashed password in the response.
	//user.Password = ""
	context.JSON(http.StatusCreated, gin.H{"message": "User created", "status": "success", "code": 200, "user": user})
	log.Println("signup")
}

func login(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse user"})
		return
	}
	log.Println("Before validate")
	if err := user.ValidateUser(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "User credentials is not valid"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User Login successfully", "status": "success", "code": 200, "user": user})
	log.Println("login")
}
