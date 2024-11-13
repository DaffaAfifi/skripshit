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

func GetAssistanceTools(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	assistanceTools, err := service.GetAssistanceTools(id, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, assistanceTools, "success get assistance tools", c)
}

func CreateAssistanceTools(c *gin.Context, db *sql.DB) {
	var request model.CreateAssistanceToolsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validation.ValidateCreateAssistanceTools(request); err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	err := service.CreateAssistanceTools(request, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, nil, "success create assistance tools", c)
}
