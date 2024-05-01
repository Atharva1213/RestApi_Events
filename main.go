package main

import (
	"Server_main/router"
	"Server_main/database"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	err := database.ConnectToDB()
    if err != nil {
        log.Println("Error connecting to database:", err)
        return
    }
	PORT := os.Getenv("PORT")
	server := gin.Default()
	server.GET("/events", router.HandlingGetEvents)
	server.POST("/eventPost", router.HandlingPostEvents)
	server.POST("/eventId", router.HandlingPostIdEvents)
	server.DELETE("/eventId", router.HandlingDeleteIdEvents)

	server.Run(PORT)
	log.Println("Server listen on " + PORT)
}
