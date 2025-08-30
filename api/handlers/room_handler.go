package handlers

import (
	"api/configs"
	"api/database/redis"
	repo "api/repositories"
	"api/types/structs"
	req "api/types/structs/requests"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	res "api/types/structs/responses"
	"api/utils/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoomIndex(c *gin.Context) {
	cached, err := redis.GetKey("rooms:all")
	if err == nil {
		var rooms []res.RoomWithDormitoryResponse
		json.Unmarshal([]byte(cached), &rooms)

		c.JSON(http.StatusOK, structs.Payload{
			Message: "Rooms retrieved successfully (from cache)",
			Error:   nil,
			Data:    rooms,
		})
		return
	}

	rooms, err := repo.GetAllRoomsWithDormitory(configs.DB)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	jsonData, err := json.Marshal(rooms)
	if err != nil { panic(err) }
	redis.SetKey("rooms:all", string(jsonData), 60)

	c.JSON(http.StatusOK, structs.Payload{
		Message: "Room retrieved successfully",
		Error:   nil,
		Data:    rooms,
	})
}

func RoomStore(c *gin.Context) {
	var roomRequest req.RoomRequest

	err := c.BindJSON(&roomRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, structs.Payload{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	v := validator.New()
	v.Required(roomRequest.DormitoryID, "dormitory_id")
	v.Required(roomRequest.RoomNumber, "room_number")
	if v.Errors() {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, structs.Payload{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    v,
		})
		return
	}

	user, err := repo.CreateRoom(configs.DB, roomRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	redis.DelKey("rooms:all")

	c.JSON(http.StatusOK, structs.Payload{
		Message: "Room inserted successfully",
		Error:   nil,
		Data:    user,
	})
}

func RoomFind(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cached, err := redis.GetKey(fmt.Sprintf("room:%d", id))
	if err == nil {
		var room res.RoomResponse
		json.Unmarshal([]byte(cached), &room)

		c.JSON(http.StatusOK, structs.Payload{
			Message: fmt.Sprintf("Room with id %d successfully found (from cache)", id),
			Error:   nil,
			Data:    room,
		})
		return
	}

	room, err := repo.GetRoomByID(configs.DB, id)
	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Room with id %d not found", id),
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

	jsonData, err := json.Marshal(room)
	if err != nil { panic(err) }
	redis.SetKey(fmt.Sprintf("room:%d", id), string(jsonData), 60)

	c.JSON(http.StatusOK, structs.Payload{
		Message: fmt.Sprintf("Room with id %d successfully found", id),
		Error:   nil,
		Data:    room,
	})
}

func RoomUpdate(c *gin.Context) {
	var roomRequest req.RoomRequest
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&roomRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, structs.Payload{
			Message: "Invalid JSON data",
			Error:   "Bad Request",
			Data:    nil,
		})
		return
	}

	v := validator.New()
	v.Required(roomRequest.DormitoryID, "dormitory_id")
	v.Required(roomRequest.RoomNumber, "room_number")
	v.Required(roomRequest.Status, "status")
	if v.Errors() {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, structs.Payload{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    v,
		})
		return
	}

	room, err := repo.UpdateRoom(configs.DB, id, roomRequest)
	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Room with id %d not found", id),
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

	redis.DelKey("rooms:all")
	redis.DelKey(fmt.Sprintf("room:%d", id))

	c.JSON(http.StatusOK, structs.Payload{
		Message: fmt.Sprintf("Room with id %d successfully updated", id),
		Error:   nil,
		Data:    room,
	})
}

func RoomDestroy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := repo.DeleteRoom(configs.DB, id)
	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Room with id %d not found", id),
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

	redis.DelKey("rooms:all")
	redis.DelKey(fmt.Sprintf("room:%d", id))

	c.JSON(http.StatusOK, structs.Payload{
		Message: fmt.Sprintf("Room with id %d successfully deleted", id),
		Error:   nil,
		Data:    nil,
	})
}

func RoomByDormitory(c *gin.Context) {
	dormitoryId, _ := strconv.Atoi(c.Param("id"))

	rooms, err := repo.GetRoomsByDormitoryID(configs.DB, dormitoryId)
	if len(rooms) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Rooms with dormitory id %d not found", dormitoryId),
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

	c.JSON(http.StatusOK, structs.Payload{
		Message: fmt.Sprintf("Rooms with dormitory id %d successfully found", dormitoryId),
		Error:   nil,
		Data:    rooms,
	})
}