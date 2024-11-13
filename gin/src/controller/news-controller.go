package controller

import (
	"database/sql"
	"gin-project/src/response"
	"gin-project/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNewsComments(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	news, err := service.GetNewsComments(id, db)
	if err != nil {
		if responseErr, ok := err.(*response.ResponseError); ok {
			c.JSON(responseErr.Status, gin.H{"error": responseErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	response.Response(200, news, "success get news comment", c)
}
