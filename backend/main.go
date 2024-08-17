package main

import (
	"example.com/online-store/db"
	"example.com/online-store/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8090")
}
