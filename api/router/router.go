package router

import (
	"api/handlers"
	"api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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

	api.GET("/dormitories/:id/rooms", handlers.RoomByDormitory)

	api.GET("/rooms", handlers.RoomIndex)
	api.POST("/rooms", handlers.RoomStore)
	api.GET("/rooms/:id", handlers.RoomFind)
	api.PUT("/rooms/:id", handlers.RoomUpdate)
	api.DELETE("/rooms/:id", handlers.RoomDestroy)

	api.GET("/rentals", handlers.RentalIndex)
	api.POST("/rentals", middlewares.AuthMiddleware, handlers.RentalStore)
	api.PATCH("/rentals/:id/status", middlewares.AuthMiddleware, handlers.RentalStatus)
	api.PATCH("/rentals/:id/duration", middlewares.AuthMiddleware, handlers.RentalAddDuration)
	api.GET("/rentals/me", middlewares.AuthMiddleware, handlers.RentalByAuthenticated)

	api.GET("/transactions", handlers.TransactionIndex)
	api.POST("/transactions/midtrans", handlers.TransactionMidtrans)
	api.POST("/transactions", handlers.TransactionStore)
	api.POST("/transactions/:id/proof", handlers.TransactionAttachProof)
	api.PATCH("/transactions/:id/status", handlers.TransactionStatus)
	api.GET("/transactions/:id", handlers.TransactionFind)

	return router
}
