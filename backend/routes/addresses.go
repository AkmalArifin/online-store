package routes

import (
	"net/http"
	"strconv"

	"example.com/online-store/models"
	"github.com/gin-gonic/gin"
)

func getAddresses(c *gin.Context) {
	addresses, err := models.GetAllAddresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get data from database. Please try again!"})
		return
	}
	c.JSON(http.StatusOK, addresses)
}

func getAdresss(c *gin.Context) {
	addressID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	var address models.Address
	address, err = models.GetAddressById(addressID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data."})
		return
	}

	c.JSON(http.StatusOK, address)
}

func getAddressesByUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	var addresses []models.Address
	addresses, err = models.GetAddressByUserId(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data."})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func createAddress(c *gin.Context) {
	var address models.Address
	err := c.ShouldBindJSON(&address)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	address.UserID = c.GetInt64("user_id")
	err = address.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse create data."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Address created!", "address": address})
}

func updateAddress(c *gin.Context) {
	addressID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	userID := c.GetInt64("user_id")
	address, err := models.GetAddressById(addressID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data."})
		return
	}

	if address.UserID != userID {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Not authorized to update address."})
		return
	}

	var updatedAddress models.Address
	err = c.ShouldBindJSON(&updatedAddress)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedAddress.ID = addressID
	err = updatedAddress.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update data."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been updated"})
}

func deleteAddress(c *gin.Context) {
	addressID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse parameter input."})
		return
	}

	userID := c.GetInt64("user_id")
	address, err := models.GetAddressById(addressID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse fetch data."})
		return
	}

	if address.UserID != userID {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Not authorized to delete address."})
		return
	}

	address.ID = addressID
	err = address.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been deleted"})
}
