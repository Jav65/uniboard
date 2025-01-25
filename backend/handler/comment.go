package handler

import (
	"net/http"
	"strconv"

	"backend/comment"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service comment.Service
}

func NewCommentHandler(service comment.Service) *CommentHandler {
	return &CommentHandler{service}
}

func (handler *CommentHandler) GetCommentByThreadIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	comments, err := handler.service.GetAllByThreadID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func (handler *CommentHandler) GetAllCommentsHandler(c *gin.Context) {
	comments, err := handler.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func (handler *CommentHandler) CreateCommentHandler(c *gin.Context) {
	var commentInput comment.CommentRequest

	err := c.ShouldBindJSON(&commentInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	comment, err := handler.service.Create(commentInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}
