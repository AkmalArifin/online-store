package main

import (
	"os"

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

	mysqlAccount := os.Getenv("MYSQL_ACCOUNT")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")

	db.InitDB(mysqlAccount, mysqlPassword)
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8090")
}
