package handlers

import (
	"api/configs"
	repo "api/repositories"
	"api/types/structs"
	req "api/types/structs/requests"
	"api/utils"
	"api/utils/validator"
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DormitoryPreviewIndex(c *gin.Context) {
	dormitoryId, _ := strconv.Atoi(c.Param("id"))

	dormitoryPreviews, err := repo.GetDormitoryPreviewsByDormitoryID(configs.DB, dormitoryId)
	if len(dormitoryPreviews) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Dormitory Previews with dormitory id %d not found", dormitoryId),
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
		Message: fmt.Sprintf("Dormitory Previews with dormitory id %d successfully found", dormitoryId),
		Error:   nil,
		Data:    dormitoryPreviews,
	})
}

func DormitoryPreviewStore(c *gin.Context) {
	var dormitoryPreviewRequest req.DormitoryPreviewRequest

	var v = validator.New()

	dormitoryId, err := strconv.Atoi(c.PostForm("dormitory_id"))
	if err != nil {
		v["dormitory_id"] = "The dormitory_id field is required."
	}

	preview, err := c.FormFile("preview")
	if err != nil {
		v["preview"] = "The preview field is required."
	}

	if v.Errors() {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, structs.Payload{
			Message: "Validation error",
			Error:   "Unprocessable Entity",
			Data:    v,
		})
		return
	}

	filename := utils.RandomString() + filepath.Ext(preview.Filename)
	err = c.SaveUploadedFile(preview, "./public/uploads/" + filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
			})
		return
	}

	dormitoryPreviewRequest.DomitoryID = dormitoryId
	dormitoryPreviewRequest.Url = fmt.Sprintf("http://localhost:8080/uploads/%s", filename)

	dormitoryPreview, err := repo.CreateDormitoryPreview(configs.DB, dormitoryPreviewRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.Payload{
			Message: "Internal server error",
			Error:   "Internal Server Error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, structs.Payload{
		Message: "Dormitory Preview inserted successfully",
		Error:   nil,
		Data:    dormitoryPreview,
	})
}

func DormitoryPreviewDestroy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := repo.DeleteDormitoryPreview(configs.DB, id)
	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Payload{
			Message: fmt.Sprintf("Dormitory Preview with id %d not found", id),
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
		Message: fmt.Sprintf("Dormitory Preview with id %d successfully deleted", id),
		Error:   nil,
		Data:    nil,
	})
}