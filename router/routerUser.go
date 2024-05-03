package router

import (
	"Server_main/user"
	"net/http"
	"github.com/gin-gonic/gin"
)

func handlingPostUser(context *gin.Context) {
	// Bind JSON data from request body to Event struct
	var user user.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User Register successfully"})
}


func handlingloginUser(context *gin.Context) {
	var user user.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens,err := user.Login()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successfully ...","tokens":tokens})
}
