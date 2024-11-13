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

func GetUsers(c *gin.Context, db *sql.DB) {
	users, err := service.GetUsers(db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, users, "success get users", c)
}

func GetUserById(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	user, err := service.GetUserById(id, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, user, "success get user by id", c)
}

func CreateUser(c *gin.Context, db *sql.DB) {
	var request model.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validation.ValidateCreateUser(request); err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	err := service.CreateUser(request, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	response.Response(200, nil, "User created successfully", c)
}

func UpdateUser(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var request model.UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validation.ValidateUpdateUser(request); err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	err := service.UpdateUser(id, request, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, nil, "success update user", c)
}

func GetSavedNews(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	savedNews, err := service.GetSavedNews(id, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, savedNews, "success get saved news", c)
}

func GetUserSavedNewsComment(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	savedNews, err := service.GetUserSavedNewsComment(id, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, savedNews, "success get saved news comment", c)
}

func GetUserFacilities(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	facilities, err := service.GetUserFacilities(id, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, facilities, "success get facilities", c)
}
