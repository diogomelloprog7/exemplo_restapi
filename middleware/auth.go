package middleware

import (
	"net/http"
	"restapi/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"-------": "Nao autorizado"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"-------": "Nao autorizado"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
