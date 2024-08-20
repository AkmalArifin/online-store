package main

import (
	"example.com/online-store/db"
	"example.com/online-store/middlewares"
	"example.com/online-store/routes"
	"github.com/gin-contrib/cors"
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

	corsConfig := middlewares.SetCORS()
	r.Use(cors.New(corsConfig))

	routes.RegisterRoutes(r)

	r.Run(":8090")
}
