package routes

import (
	"database/sql"
	"gin-project/src/controller"
	"gin-project/src/middleware"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(router *gin.Engine, db *sql.DB) {
	router.Use(middleware.AuthMiddleware(db))
	public := router.Group("/api")

	// Users
	public.GET("/users", func(c *gin.Context) {
		controller.GetUsers(c, db)
	})
	public.GET("/users/:id", func(c *gin.Context) {
		controller.GetUserById(c, db)
	})
	public.POST("/users", func(c *gin.Context) {
		controller.CreateUser(c, db)
	})
	public.PUT("/users/:id", func(c *gin.Context) {
		controller.UpdateUser(c, db)
	})
	public.GET("/users/saved-news/:id", func(c *gin.Context) {
		controller.GetSavedNews(c, db)
	})
	public.GET("/users/saved-news/comment/:id", func(c *gin.Context) {
		controller.GetUserSavedNewsComment(c, db)
	})
	public.GET("/users/facilities/:id", func(c *gin.Context) {
		controller.GetUserFacilities(c, db)
	})

	// news
	public.GET("/news/:id", func(c *gin.Context) {
		controller.GetNewsComments(c, db)
	})

	// assistance
	public.GET("/assistance/:id", func(c *gin.Context) {
		controller.GetAssistanceTools(c, db)
	})
	public.POST("/assistance-tools", func(c *gin.Context) {
		controller.CreateAssistanceTools(c, db)
	})

	// auth
	public.POST("/logout", func(c *gin.Context) {
		controller.Logout(c, db)
	})
}
