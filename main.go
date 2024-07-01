package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/PING", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "Abuzar G from Lahore",
		})
	})

	router.GET("/abx/:id/:buyerID", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
		buyerID := c.Param("buyerID")
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"buyerID": buyerID,
		})

	})

	router.POST("/test", func(c *gin.Context) {
		// Email and Pass

		type TestRequest struct {
			Email string `json:"email" binding:"required"`
			Pass  string `json:"pass"`
		}
		var req TestRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// c.BindJSON(&req)
		c.JSON(http.StatusOK, gin.H{
			"email": req.Email,
			"pass":  req.Pass,
		})

	})

	router.Run(":8400")
}
