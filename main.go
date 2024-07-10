package main

import (
	"GinMod/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	UserController := &controllers.UserController{}
	UserController.InitUserControllerRoutes(router)

	router.Run(":8000")
}
