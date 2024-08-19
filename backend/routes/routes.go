package routes

import (
	"example.com/online-store/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// Group
	auth := r.Group("/")
	auth.Use(middlewares.Authenticate)
	// User
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.POST("/login", login)
	// r.POST("/signup")

	// Address
	r.GET("/addresses", getAddresses)
	r.GET("/addresses/:id", getAdresss)
	r.GET("/addresses/user/:id", getAddressesByUser)
	auth.POST("/addresses", createAddress)
	auth.PUT("/addresses/:id", updateAddress)
	auth.DELETE("/addresses/:id", deleteAddress)
}
