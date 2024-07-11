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
		users, err := u.UserService.GetUserService()

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return

		}
		c.JSON(200, gin.H{

			"users": users,
		})
	}

}

func (u *UserController) CreateUser() gin.HandlerFunc {
	type UserBody struct {
		Name   string `json:"name"`
		Status bool   `json:"status"`
	}

	return func(c *gin.Context) {
		var userBody UserBody
		if err := c.BindJSON(&userBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := u.UserService.CreateUserService(userBody.Name, userBody.Status)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}
