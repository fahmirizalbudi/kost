package handlers

import (
	"api/configs"
	repo "api/repositories"
	"api/types/structs"
	req "api/types/structs/requests"
	// res "api/types/structs/responses"
	"api/utils/validate"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	users, err := repo.GetAllUsers(configs.DB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Payload{
		Message: "Users retrieved successfully",
		Error:   nil,
		Data: users,
	})
}

func UserStore(c *gin.Context) {
	var userRequest req.UserRequest

	err := c.BindJSON(&userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, structs.Payload{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	validations := map[string]string{}
	validate.Required(validations, userRequest.Name, "name")
	validate.Required(validations, userRequest.Email, "email")
	validate.Required(validations, userRequest.Password, "password")
	validate.Required(validations, userRequest.Role, "role")
	validate.Required(validations, userRequest.Phone, "phone")
	validate.Required(validations, userRequest.Address, "address")
	if len(validations) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, structs.Payload{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    gin.H{
				"validations": validations,
			},
		})
		return
	}

	user, err := repo.CreateUser(configs.DB, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Payload{
		Message: "User inserted successfully",
		Error:   nil,
		Data: user,
	})
}
