package router

import (
	"Server_main/events"
	"net/http"
	"github.com/gin-gonic/gin"
)

func handlingGetEvents(context *gin.Context) {
	result, err := events.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, result)
}
func handlingPostEvents(context *gin.Context) {
	// Bind JSON data from request body to Event struct
	var event events.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    
	userID,_:= context.Get("userID")
	err := event.Save(userID.(int64))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}
func handlingPostIdEvents(context *gin.Context) {
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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": result})
}
func handlingDeleteIdEvents(context *gin.Context) {
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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID,_:= context.Get("userID")
	if result.UserId != userID {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Owner Only Manipulated the Event"})
		return
	}
	err = events.DeleteEvent(currentReqBody.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleting successfully","event":result})
}
func handlingUpdateIdEvents(context *gin.Context) {

	var event events.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID,_:= context.Get("userID")
	if event.UserId != userID {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Owner Only Manipulated the Event"})
		return
	}
	err := event.UpdatedEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message":"Event Updated Successfully"})
}


