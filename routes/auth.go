package routes

import (
	"api/models"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var user models.User
	utils.BindAndRespondBadRequest(&user, context)

	err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not register the user"})
		return
	}

	token := utils.GenerateTokenbAndRespondInternal(user.Email, user.Id, context)

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(context *gin.Context) {
	var user models.User
	utils.BindAndRespondBadRequest(&user, context)

	err := user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token := utils.GenerateTokenbAndRespondInternal(user.Email, user.Id, context)

	context.JSON(http.StatusOK, gin.H{"token": token})
}
