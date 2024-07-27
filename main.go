package main

import (
	"api/db"
	"api/routes"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.Use(cors.Default())

	routes.RegisterRoutes(server)

	posrt := os.Getenv("PORT")

	if posrt == "" {
		posrt = "8080"
	}

	server.Run(fmt.Sprintf("0.0.0.0:%s", posrt))
}
