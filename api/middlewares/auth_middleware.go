package middlewares

import (
	"api/helpers"
	"api/types/structs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusOK, structs.Payload{
				Message: "Invalid or expired token",
				Error:   "Unauthorized",
				Data:    nil,
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { 
			c.AbortWithStatusJSON(http.StatusOK, structs.Payload{
				Message: "Invalid or expired token",
				Error:   "Unauthorized",
				Data:    nil,
			})
			return
		}

		claims, err := helpers.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, structs.Payload{
				Message: "Invalid or expired token",
				Error:   "Unauthorized",
				Data:    nil,
			})
			return
		}

		c.Set("claims", claims.(*helpers.Claims))
		c.Next()
}