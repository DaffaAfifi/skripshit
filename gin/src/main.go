package main

import (
	"fmt"
	"gin-project/src/application"
	"gin-project/src/validation"
)

func main() {
	application.InitDB()
	validation.InitValidator()
	port := 8080
	router := application.SetupRouter(application.DB)
	application.InitLogger()
	application.Logger.Info("App started and running on port " + fmt.Sprintf("%d", port))
	application.SyncLogger()
	router.Run(fmt.Sprintf(":%d", port))
}
