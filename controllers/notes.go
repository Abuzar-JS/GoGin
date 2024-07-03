package controllers

import (
	"GinMod/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	notes.PUT("/", n.PutNotes())
	notes.DELETE("/", n.DeleteNotes())

}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(200, gin.H{
			"notes": n.notesService.GetNotesService(),
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
			"notes": n.notesService.PostNotesService(),
		})
	}
}

func (n *NotesController) DeleteNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.notesService.GetNotesService(),
		})
	}
}
