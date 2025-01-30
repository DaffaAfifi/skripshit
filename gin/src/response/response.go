package response

import "github.com/gin-gonic/gin"

// Response mengirimkan response JSON dengan status, data, pesan, dan metadata ke client.
func Response(statusCode int, data interface{}, message string, c *gin.Context) {
	c.JSON(statusCode, gin.H{
		"payload": data,
		"message": message,
		"metadata": gin.H{
			"prev":    "",
			"next":    "",
			"current": "",
		},
	})
}
