package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

type NoteRoutes struct {
	noteHandler handlers.NoteHandler
}

func NewNoteRoutes(noteHandler handlers.NoteHandler) NoteRoutes {
	return NoteRoutes{noteHandler}
}

func (nr *NoteRoutes) SetNoteRoutes(router *gin.RouterGroup) {
	notes := router.Group("/notes")

	notes.GET("/", nr.noteHandler.GetAllNote)
	notes.GET("/:id", nr.noteHandler.GetNote)
	notes.GET("/:id/user", nr.noteHandler.GetNoteByUser)
	notes.POST("/", nr.noteHandler.CreateNote)
	notes.POST("/:id", nr.noteHandler.UpdateNote)
	notes.DELETE("/:id", nr.noteHandler.DeleteNote)

}
