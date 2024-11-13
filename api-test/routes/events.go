package routes

import (
	"net/http"
	"restapi/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mmensagem": "Not PARSE ID!"})
	}
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mmensagem": "erro!"})
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mmensagem": "Nao pode ser cadastrado este evento"})
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi converter "})
	}
	userId := context.GetInt64("userId")
	event.ID = 1

	event.UserID = 1
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mmensagem": "WTF"})
	}
	context.JSON(http.StatusCreated, gin.H{"-------": "Evento criado"})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi Converter os dados "})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi possivel atualizar a tabela "})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"": "NAO AUTORIZADO POR GIN"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi possivel atualizar a tabela "})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi possivel atualizar o evento "})
		return
	}
	context.JSON(http.StatusOK, gin.H{"------": "Atualizado !"})
}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi Converter os dados "})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"": "NAO AUTORIZADO POR GIN"})
		return
	}

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi possivel atualizar a tabela "})
		return

	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"-------": "Nao foi possivel deletar o item "})
		return
	}
	context.JSON(http.StatusOK, gin.H{"-------": "ITEM DELETADO"})

}
