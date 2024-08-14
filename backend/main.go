package main

import (
	"net/http"

	"example.com/online-store/db"
	"example.com/online-store/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/", getEvents)

	// User
	r.GET("/users", getUsers)
	r.POST("/users", createUser)

	r.Run(":8090")
}

func getUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get data from database. Please try again!", "error": err.Error()})
	}
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse create data.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
