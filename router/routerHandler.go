package router

import (
	"github.com/gin-gonic/gin"
	"Server_main/middleware" 
)

func RouterHandler(server *gin.Engine) {
	authRoutes := server.Group("")
	authRoutes.Use(middleware.Authenticate)
	{
		authRoutes.POST("/events", handlingPostEvents)
		authRoutes.DELETE("/eventId", handlingDeleteIdEvents)
		authRoutes.PUT("/eventId", handlingUpdateIdEvents)
	}

	server.GET("/events", handlingGetEvents)
	server.POST("/eventId", handlingPostIdEvents)
	server.POST("/register", handlingPostUser)
	server.POST("/login", handlingloginUser)
}
