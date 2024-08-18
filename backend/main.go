package main

import (
	"example.com/online-store/db"
	"example.com/online-store/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(".env cannot be load.")
	}

	db.InitDB()
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8090")
}
