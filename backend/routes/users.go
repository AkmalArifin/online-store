package routes

import (
	"net/http"
	"strconv"
	"time"

	"example.com/online-store/models"
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get data from database. Please try again!", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	var user models.User
	user, err = models.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
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

func updateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	_, err = models.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data.", "error": err.Error()})
		return
	}

	var updatedUser models.User
	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	updatedUser.ID = userId
	err = updatedUser.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update data.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been updated"})
}

func deleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	var user models.User
	user, err = models.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data.", "error": err.Error()})
		return
	}

	user.DeletedAt.SetValue(time.Now())
	user.ID = userId

	err = user.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been deleted"})
}
