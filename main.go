package main

import (
	"GinMod/controllers"
	internal "GinMod/internal/database"
	"GinMod/services"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db := internal.DatabaseConnection()
	if db == nil {
		return
	}

	userService := &services.UserService{}
	userService.InitService(db)

	userController := &controllers.UserController{UserService: *userService}
	userController.InitUserControllerRoutes(router)

	router.Run(":8000")
}
