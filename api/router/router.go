package router

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.POST("/auth/register", handlers.Register)
	api.POST("/auth/login", handlers.Login)
	api.POST("/auth/me", handlers.Me)
	api.POST("/auth/logout", handlers.Logout)

	api.GET("/users", handlers.UserIndex)
	api.POST("/users", handlers.UserStore)
	api.GET("/users/:id", handlers.UserFind)
	api.PUT("/users/:id", handlers.UserUpdate)
	api.DELETE("/users/:id", handlers.UserDestroy)

	api.GET("/dormitories", handlers.DormitoryIndex)
	api.POST("/dormitories", handlers.DormitoryStore)
	api.GET("/dormitories/:id", handlers.DormitoryFind)
	api.PUT("/dormitories/:id", handlers.DormitoryUpdate)
	api.DELETE("/dormitories/:id", handlers.DormitoryDestroy)

	return router
}