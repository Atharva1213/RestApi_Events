package middleware

import (
	"net/http"
	"Server_main/utilty"
	"github.com/gin-gonic/gin"
)


func Authenticate(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
            return
        }

        // Verify the token and extract user ID
        userID, err := utilty.VerifyToken(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }
       // Add user ID to request context
        c.Set("userID", userID)

        // Call the next handler
        c.Next()
}