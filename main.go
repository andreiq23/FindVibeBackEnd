package main

import (
	"api/db"
	"api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.Use(cors.Default())

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
