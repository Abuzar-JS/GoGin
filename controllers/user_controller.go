package controllers

import (
	"GinMod/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func (u *UserController) InitUserControllerRoutes(router *gin.Engine) {
	users := router.Group("/user")
	users.GET("/", u.GetUser())
	users.POST("/", u.CreateUser())
}

func (u *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"User": u.UserService.GetUserService(),
		})
	}
}

func (u *UserController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"User": "Create User Request",
		})
	}
}
