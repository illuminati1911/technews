package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/illuminati1911/technews/models"
	"github.com/illuminati1911/technews/service-auth/auth"
)

type AuthHTTPHandler struct {
	as auth.Service
}

type userParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// NewAuthHTTPHandler creates REST API endpoints for given auth service
func NewAuthHTTPHandler(as auth.Service, router *gin.Engine) *AuthHTTPHandler {
	handler := &AuthHTTPHandler{as}
	router.Use(gin.Recovery())
	router.POST("/login", handler.login)
	router.POST("/create", handler.createUser)
	return handler
}

func (a *AuthHTTPHandler) login(c *gin.Context) {
	var json userParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	token, err := a.as.Login(json.Username, json.Password)
	if err != nil {
		if te, ok := err.(*models.TNError); ok {
			c.JSON(te.HttpCode, gin.H{"error": te.Message})
			return
		}
		c.JSON(models.ErrGeneralServerError.Code, gin.H{
			"error": models.ErrGeneralServerError.Message,
		})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

func (a *AuthHTTPHandler) createUser(c *gin.Context) {
	var json userParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	_, err := a.as.CreateUser(json.Username, json.Password)
	if err != nil {
		if te, ok := err.(*models.TNError); ok {
			c.JSON(te.HttpCode, gin.H{"error": te.Message})
			return
		}
		c.JSON(models.ErrGeneralServerError.Code, gin.H{
			"error": models.ErrGeneralServerError.Message,
		})
		return
	}
	c.Status(200)
}
