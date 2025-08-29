package handlers

import (
	"api/configs"
	"api/database/redis"
	repo "api/repositories"
	"api/types/structs"
	req "api/types/structs/requests"
	"encoding/json"

	"fmt"
	"database/sql"
	"strconv"

	res "api/types/structs/responses"
	"api/utils/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DormitoryIndex(c *gin.Context) {
	cached, err := redis.GetKey("dormitories:all")
	if err == nil {
		var dormitories []res.DormitoryResponse
		json.Unmarshal([]byte(cached), &dormitories)

		c.JSON(http.StatusOK, structs.Payload{
			Message: "Dormitories retrieved successfully (from cache)",
			Error:   nil,
			Data:    dormitories,
		})
		return
	}

	dormitories, err := repo.GetAllDormitories(configs.DB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	jsonData, err := json.Marshal(dormitories)
	if err != nil { panic(err) }
	redis.SetKey("dormitories:all", string(jsonData), 60)

	c.JSON(http.StatusOK, structs.Payload{
		Message: "Dormitories retrieved successfully",
		Error:   nil,
		Data:    dormitories,
	})
}

func DormitoryStore(c *gin.Context) {
	var dormitoryRequest req.DormitoryRequest

	err := c.BindJSON(&dormitoryRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, structs.Payload{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	v := validator.New()
	v.Required(dormitoryRequest.Name, "name")
	v.Required(dormitoryRequest.Address, "address")
	v.Required(dormitoryRequest.Description, "description")
	v.Required(dormitoryRequest.Price, "price")
	v.Required(dormitoryRequest.Facilities, "facilities")
	v.Required(dormitoryRequest.GoogleMaps, "google_maps")
	if v.Errors() {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, structs.Payload{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    v,
		})
		return
	}

	dormitory, err := repo.CreateDormitory(configs.DB, dormitoryRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	redis.DelKey("dormitories:all")

	c.JSON(http.StatusOK, structs.Payload{
		Message: "Dormitory inserted successfully",
		Error:   nil,
		Data:    dormitory,
	})
}

func DormitoryFind(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cached, err := redis.GetKey(fmt.Sprintf("dormitory:%d", id))
	if err == nil {
		var dormitory res.DormitoryResponse
		json.Unmarshal([]byte(cached), &dormitory)

		c.JSON(http.StatusOK, structs.Payload{
			Message: fmt.Sprintf("Dormitory with id %d successfully found (from cache)", id),
			Error:   nil,
			Data:    dormitory,
		})
		return
	}

	dormitory, err := repo.FindDormitory(configs.DB, id)
	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Dormitory with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	jsonData, err := json.Marshal(dormitory)
	if err != nil { panic(err) }
	redis.SetKey(fmt.Sprintf("dormitory:%d", id), string(jsonData), 60)

	c.JSON(http.StatusOK, structs.Payload{
		Message: fmt.Sprintf("Dormitory with id %d successfully found", id),
		Error:   nil,
		Data:    dormitory,
	})
}

func DormitoryUpdate(c *gin.Context) {
	var dormitoryRequest req.DormitoryRequest
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&dormitoryRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, structs.Payload{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	v := validator.New()
	v.Required(dormitoryRequest.Name, "name")
	v.Required(dormitoryRequest.Address, "address")
	v.Required(dormitoryRequest.Description, "description")
	v.Required(dormitoryRequest.Price, "price")
	v.Required(dormitoryRequest.Facilities, "facilities")
	v.Required(dormitoryRequest.GoogleMaps, "google_maps")
	if v.Errors() {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, structs.Payload{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    v,
		})
		return
	}

	dormitory, err := repo.UpdateDormitory(configs.DB, id, dormitoryRequest)
	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Dormitory with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	redis.DelKey("dormitories:all")
	redis.DelKey(fmt.Sprintf("dormitory:%d", id))

	c.JSON(http.StatusOK, structs.Payload{
		Message: fmt.Sprintf("Dormitory with id %d successfully updated", id),
		Error:   nil,
		Data:    dormitory,
	})
}

func DormitoryDestroy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := repo.DeleteDormitory(configs.DB, id)
	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Dormitory with id %d not found", id),
			Error:   "Not Found",
			Data:    nil,
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	redis.DelKey("dormitories:all")
	redis.DelKey(fmt.Sprintf("dormitory:%d", id))

	c.JSON(http.StatusOK, structs.Payload{
		Message: fmt.Sprintf("Dormitory with id %d successfully deleted", id),
		Error:   nil,
		Data:    nil,
	})
}