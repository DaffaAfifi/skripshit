package middleware

import (
	"gin-project/src/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors.Last()
		if errs != nil {
			if responseError, ok := errs.Err.(*response.ResponseError); ok {
				c.JSON(responseError.Status, gin.H{
					"errors": responseError.Message,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"errors": errs.Error(),
				})
			}
			c.Abort()
		}
	}
}
