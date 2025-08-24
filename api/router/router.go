package router

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/users", handlers.UserIndex)

	return router
}