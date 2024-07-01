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

	meetings.GET("/:id", mr.meetingHandler.GetMeeting)

}
