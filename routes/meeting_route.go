package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

type MeetingRoutes struct {
	meetingHandler handlers.MeetingHandler
}

func NewMeetingRoutes(meetingHandler handlers.MeetingHandler) MeetingRoutes {
	return MeetingRoutes{meetingHandler}
}

func (mr *MeetingRoutes) SetMeetingRoutes(router *gin.RouterGroup) {
	meetings := router.Group("/meetings")

	// todos.GET("/", tr.todoHandler.GetAllTodo)
	meetings.GET("/:id", mr.meetingHandler.GetMeeting)
	meetings.POST("/", mr.meetingHandler.CreateMeeting)
	// todos.PATCH("/:id", tr.todoHandler.UpdateTodo)
	// todos.DELETE("/:id", tr.todoHandler.DeleteTodo)
}
