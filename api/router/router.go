package router

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	router.Static("/uploads", "./public/uploads")

	api := router.Group("/api")

	api.POST("/auth/register", handlers.Register)
	api.POST("/auth/login", handlers.Login)
	api.POST("/auth/me", handlers.Me)

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

	api.GET("/dormitories/previews", handlers.DormitoryAttachPreviews)
	api.GET("/dormitories/:id/previews", handlers.DormitoryPreviewIndex)

	api.POST("/dormitory-previews", handlers.DormitoryPreviewStore)
	api.DELETE("/dormitory-previews/:id", handlers.DormitoryPreviewDestroy)

	return router
}