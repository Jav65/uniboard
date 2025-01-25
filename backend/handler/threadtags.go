package handler

import (
	"net/http"
	"strconv"

	"backend/threadtags"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ThreadTagsHandler struct {
	service threadtags.Service
}

func NewThreadTagsHandler(service threadtags.Service) *ThreadTagsHandler {
	return &ThreadTagsHandler{service}
}

func (handler *ThreadTagsHandler) GetThreadTagsByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	threadtags, err := handler.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": threadtags,
	})
}


func (handler *ThreadTagsHandler) GetThreadTagsByTagIDHandler(c *gin.Context) {
	idStr := c.Param("tag_id")
	id, _ := strconv.Atoi(idStr)

	threadtags, err := handler.service.GetByTagID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": threadtags,
	})
}

func (handler *ThreadTagsHandler) GetAllThreadTagsHandler(c *gin.Context) {
	threadtagss, err := handler.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": threadtagss,
	})
}

func (handler *ThreadTagsHandler) CreateThreadTagsHandler(c *gin.Context) {
	var threadtagsInput threadtags.ThreadTagsRequest

	err := c.ShouldBindJSON(&threadtagsInput)
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

	threadtags, err := handler.service.Create(threadtagsInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": threadtags,
	})
}

func (handler *ThreadTagsHandler) UpdateThreadTagsHandler(c *gin.Context) {
	var threadtagsInput threadtags.ThreadTagsRequest

	err := c.ShouldBindJSON(&threadtagsInput)
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

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	_, err = handler.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	threadtags, err := handler.service.Update(id, threadtagsInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": threadtags,
	})
}

func (handler *ThreadTagsHandler) DeleteThreadTagsHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	threadtags, err := handler.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	_, err = handler.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": threadtags,
	})
}
