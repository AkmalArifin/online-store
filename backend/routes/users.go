package routes

import (
	"net/http"
	"strconv"

	"example.com/online-store/models"
	"example.com/online-store/utils"
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get data from database. Please try again!"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	var user models.User
	user, err = models.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data."})
		return
	}

	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse create data."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func updateUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	_, err = models.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data."})
		return
	}

	var updatedUser models.User
	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedUser.ID = userID
	err = updatedUser.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update data."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been updated"})
}

func deleteUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	var user models.User
	user, err = models.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data."})
		return
	}

	user.ID = userID
	err = user.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been deleted"})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorized account."})
		return
	}

	var token string
	token, err = utils.GenerateToken(user.Email.String, user.ID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorized account."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login succes!", "token": token})
}
