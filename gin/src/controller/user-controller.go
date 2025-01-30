package controller

import (
	"database/sql"
	"gin-project/src/model"
	"gin-project/src/response"
	"gin-project/src/service"
	"gin-project/src/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers meng-handle permintaan untuk mendapatkan daftar semua pengguna dari database.
func GetUsers(c *gin.Context, db *sql.DB) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "100")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}

	users, err := service.GetUsers(db, page, limit)
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

// GetUserById meng-handle permintaan untuk mendapatkan data pengguna berdasarkan ID pengguna.
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

// CreateUser meng-handle permintaan untuk membuat pengguna baru dengan validasi input dan proses pembuatan.
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

// UpdateUser meng-handle permintaan untuk memperbarui data pengguna berdasarkan ID pengguna.
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

// GetSavedNews meng-handle permintaan untuk mendapatkan berita yang disimpan oleh pengguna berdasarkan ID pengguna.
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

// GetUserSavedNewsComment meng-handle permintaan untuk mendapatkan komentar pada berita yang disimpan oleh pengguna.
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

// GetUserFacilities meng-handle permintaan untuk mendapatkan fasilitas yang dimiliki oleh pengguna berdasarkan ID pengguna.
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
