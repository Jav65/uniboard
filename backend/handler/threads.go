package handler

import (
	"net/http"
	"strconv"

	"backend/thread"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ThreadHandler struct {
	service thread.Service
}

func NewThreadHandler(service thread.Service) *ThreadHandler {
	return &ThreadHandler{service}
}

func (handler *ThreadHandler) GetThreadByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	thread, err := handler.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": thread,
	})
}

func (handler *ThreadHandler) GetSortedThreadsHandler(c *gin.Context) {
	sortBy := c.Query("sortBy")
	search := c.Query("search")

	threads, err := handler.service.GetSorted(sortBy, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": threads,
	})
}

func (handler *ThreadHandler) CreateThreadHandler(c *gin.Context) {
	var threadInput thread.ThreadRequest

	err := c.ShouldBindJSON(&threadInput)
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

	thread, err := handler.service.Create(threadInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": thread,
	})
}

func (handler *ThreadHandler) UpdateThreadHandler(c *gin.Context) {
	var threadInput thread.ThreadRequest

	err := c.ShouldBindJSON(&threadInput)
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

	thread, err := handler.service.Update(id, threadInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": thread,
	})
}

func (handler *ThreadHandler) DeleteThreadHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	thread, err := handler.service.GetByID(id)
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
		"data": thread,
	})
}
