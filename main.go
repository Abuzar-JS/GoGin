package main

import (
	"GinMod/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	notesController := controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router)

	router.Run(":8200")
}
