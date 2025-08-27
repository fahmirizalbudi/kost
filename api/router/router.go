package router

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/users", handlers.UserIndex)
	api.POST("/users", handlers.UserStore)
	api.GET("/users/:id", handlers.UserFind)
	api.PUT("/users/:id", handlers.UserUpdate)

	return router
}