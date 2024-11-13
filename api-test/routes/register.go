package routes

import (
	"net/http"
	"restapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi Converter os dados "})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"-------": "Nao foi possivel executar esta açao"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"-------": "Nao foi possivel executar esta açao"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"": "DEU BOA!"})
}

func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"-------": "Nao foi possivel executar esta açao"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"": "NICE!"})
}
