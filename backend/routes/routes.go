package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {

	// User
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.POST("/signup", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
}
