package controllers

import (
	"GinMod/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func (u *UserController) InitUserControllerRoutes(router *gin.Engine) {
	users := router.Group("/user")
	users.GET("/", u.GetUser())
	// users.GET("/:id", u.GetUserById())
	users.POST("/", u.CreateUser())
	users.PUT("/", u.UpdateUser())
	users.DELETE("/:id", u.DeleteUser())
}

func (u *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		status := c.Query("status")
		var actualStatus *bool
		if status != "" {
			aS, err := strconv.ParseBool(status)
			actualStatus = &aS
			if err != nil {
				c.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}

			users, err := u.UserService.GetUserService(actualStatus)
			if err != nil {
				c.JSON(400, gin.H{

					"message": err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"user": users,
			})
		}

	}
}

// func (u *UserController) GetUserById() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id := c.Param("id")
// 		userId, err := strconv.ParseInt(id, 10, 64)
// 		if err != nil {
// 			c.JSON(404, gin.H{})
// 		}

// 	}
// }

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

func (u *UserController) UpdateUser() gin.HandlerFunc {
	type UserBody struct {
		Name   string `json:"name"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding:"required"`
	}

	return func(c *gin.Context) {
		var userBody UserBody
		if err := c.BindJSON(&userBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := u.UserService.UpdateUserService(userBody.Name, userBody.Status, userBody.Id)
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

func (u *UserController) DeleteUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		id := c.Param("id")
		userId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err,
			})
			return
		}

		err = u.UserService.DeleteUserService(userId)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "User Deleted Successfully!",
		})
	}
}
