package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/auth/register", Register)
	server.POST("/auth/login", Login)
	server.GET("/songs", SearchSongs)
}
