package application

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"gin-project/src/middleware"
	"gin-project/src/routes"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	routes.PublicRoutes(router, db)
	routes.PrivateRoutes(router, db)
	router.Use(middleware.ErrorMiddleware())

	return router
}
