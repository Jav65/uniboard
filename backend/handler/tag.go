package handler

import (
	"net/http"
	"strconv"

	"backend/tag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TagHandler struct {
	service tag.Service
}

func NewTagHandler(service tag.Service) *TagHandler {
	return &TagHandler{service}
}

func (handler *TagHandler) GetTagByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	tag, err := handler.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tag,
	})
}

func (handler *TagHandler) GetAllNameTagsHandler(c *gin.Context) {
	tags, err := handler.service.GetAllName()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tags,
	})
}

func (handler *TagHandler) GetAllTagsHandler(c *gin.Context) {
	tags, err := handler.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tags,
	})
}

func (handler *TagHandler) CreateTagHandler(c *gin.Context) {
	var tagInput tag.TagRequest

	err := c.ShouldBindJSON(&tagInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error: %s, Condition: %s\n", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessages,
		})
		return
	}

	tag, err := handler.service.Create(tagInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tag,
	})
}