package router

import (
	"Server_main/events"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandlingGetEvents(context *gin.Context) {
	result, err := events.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, result)
}

func HandlingPostEvents(context *gin.Context) {
	// Bind JSON data from request body to Event struct
	var event events.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in  Inserting Data ..."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}
func HandlingPostIdEvents(context *gin.Context) {
	type reqBody struct {
		ID int64 `json:"id"`
	}

	var currentReqBody reqBody
	if err := context.ShouldBindJSON(&currentReqBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := events.GetEventByID(currentReqBody.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": result})
}
func HandlingDeleteIdEvents(context *gin.Context) {
	type reqBody struct {
		ID int64 `json:"id"`
	}

	var currentReqBody reqBody
	if err := context.ShouldBindJSON(&currentReqBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := events.DeleteEvent(currentReqBody.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in  Inserting Data ..."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleting successfully"})
}



