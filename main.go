package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/PING", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "Abuzar from Lahore",
		})
	})
	r.Run(":8400")
}
