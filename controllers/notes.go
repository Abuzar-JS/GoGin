package controllers

import (
	"github.com/gin-gonic/gin"
)

type NotesController struct {
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	notes.PUT("/", n.PutNotes())

}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(200, gin.H{
			"notes": "Get Request Notes",
		})
	}
}

func (n *NotesController) PutNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": "Put Request Notes",
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": "POST Request Notes",
		})
	}
}
