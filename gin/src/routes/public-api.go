package routes

import (
	"database/sql"
	"gin-project/src/controller"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(router *gin.Engine, db *sql.DB) {
	public := router.Group("/api")

	public.POST("/login", func(c *gin.Context) {
		controller.Login(c, db)
	})
}
