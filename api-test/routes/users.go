package routes

import (
	"net/http"
	"restapi/models"
	"restapi/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "WTF"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "WTF"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "WTF"})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "WTF"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"-------": "LOGIN INVALIDO"})
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"-------": "LOGIN INVALIDO"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"-----": "token", "token": token})
	return
}
