package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type TestRequest struct {
// 	Email string `json:"email" binding:"required"`
// 	Pass  string `json:"pass"`
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/PING", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"Hello": "Abuzar from Lahore",
// 		})
// 	})

// 	router.GET("/abx/:name/:ID/:Location", func(c *gin.Context) {
// 		Name := c.Param("name")
// 		ID := c.Param("ID")
// 		Location := c.Param("Location")
// 		c.JSON(http.StatusOK, gin.H{
// 			"Location": Location,
// 			"Name":     Name,
// 			"ID":       ID,
// 		})

// 	})

// 	router.POST("/test", func(c *gin.Context) {
// 		// Email and Pass

// 		var req TestRequest
// 		if err := c.BindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"email": req.Email,
// 			"pass":  req.Pass,
// 		})

// 	})

// 	router.PATCH("/test", func(c *gin.Context) {
// 		// Email and Pass

// 		type TestRequest struct {
// 			Email string `json:"email" binding:"required"`
// 			Pass  string `json:"pass"`
// 		}
// 		var req TestRequest
// 		if err := c.BindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"email": req.Email,
// 			"pass":  req.Pass,
// 		})

// 	})

// 	router.DELETE("/test/:id", func(c *gin.Context) {
// 		id := c.Param("id")

// 		c.JSON(http.StatusOK, gin.H{
// 			"id":      id,
// 			"message": "deleted!!",
// 		})

// 	})

// 	router.Run(":8405")
// }
