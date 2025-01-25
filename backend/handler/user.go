package handler

import (
	"backend/user"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{service}
}

const SecretKey = "secret"

func (handler *UserHandler) Register(c *gin.Context) {
	var userInput user.UserRequest

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	user, err := handler.service.Register(userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (handler *UserHandler) Login(c *gin.Context) {
	var userInput user.UserLogin

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	user, err := handler.service.Login(userInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err,
		})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)

	c.SetCookie(
		"jwt",
		token,
		24*3600,
		"/",
		"localhost",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"data":    user.Username,
	})
}

func (handler *UserHandler) User(c *gin.Context) {
	cookie, err := c.Cookie("jwt")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err,
		})
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err,
		})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := handler.service.CheckUser(claims.Issuer)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"id":       user.Id,
		"email":    user.Email,
	})
}

func (handler *UserHandler) Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)

	c.SetCookie(
		"jwt",
		"a",
		-1,
		"/",
		"localhost",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
