package routes

import (
	"api/models"
	"api/utils"

	"github.com/gin-gonic/gin"
)

func SearchSongs(context *gin.Context) {
	search := context.Query("search")
	search, err := utils.CleanString(search)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	songs, err := models.SearchSongs(search)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, songs)
}
