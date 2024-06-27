package middleware

import (
	"golang-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
            c.Abort()
            return
        }

        // Bearer tokenString
        tokenString = strings.TrimSpace(strings.Replace(tokenString, "Bearer", "", 1))

        claims, err := utils.VerifyToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}
