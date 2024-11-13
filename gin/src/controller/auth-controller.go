package controller

import (
	"database/sql"
	"gin-project/src/model"
	"gin-project/src/response"
	"gin-project/src/service"
	"gin-project/src/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, db *sql.DB) {
	var request model.LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validation.ValidateLogin(request); err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	token, err := service.Login(request, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, gin.H{"token": token}, "success login", c)
}

func Logout(c *gin.Context, db *sql.DB) {
	token := c.GetHeader("Authorization")
	err := service.Logout(token, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, nil, "success logout", c)
}
