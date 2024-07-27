package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindAndRespondBadRequest(model any, context *gin.Context) {
	err := context.ShouldBindBodyWithJSON(&model)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not parse json"})
		return
	}
}

func GenerateTokenbAndRespondInternal(email string, userId int64, context *gin.Context) string {
	token, err := generateToken(email, userId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return ""
	}
	return token
}
