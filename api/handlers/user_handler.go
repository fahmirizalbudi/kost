package handlers

import (
	"api/configs"
	repo "api/repositories"
	"api/types/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	users, err := repo.GetAllUsers(configs.DB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error: "Internal Server Error",
			Data: nil,
		})
	}

	c.JSON(http.StatusOK, structs.Payload{
		Message: "Users retrieved successfully",
		Error: nil,
		Data: gin.H{
			"users": users,
		},
	})
}