package handlers

import (
	"net/http"

	_ "github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

type MeetingHandler struct {
	meetingService services.MeetingService
}

func NewMeetingHandler(meetingService services.MeetingService) MeetingHandler {
	return MeetingHandler{meetingService}
}

// GetMeeting godoc
// @Tags Meeting
// @Summary Meeting 조회
// @Description Meeting 조회
// @ID GetMeeting
// @Accept  json
// @Produce  json
// @Param meetingId path string true "Meeting ID"
// @Router /meetings/{meetingId} [get]
// @Success 200 {object} dto.APIResponse[Meeting]
// @Failure 500
func (mh *MeetingHandler) GetMeeting(ctx *gin.Context) {
	meetingId := ctx.Param("id")

	meeting, err := mh.meetingService.GetMeeting(meetingId)

	if err != nil {
		// CustomError 인터페이스로 형변환이 성공하면 customErr에는 *errors.CustomError 타입의 값이 할당되고, ok 변수에는 true가 할당
		customErr, ok := err.(*errors.CustomError)
		if ok {
			statusCode := customErr.Status()
			ctx.JSON(statusCode, gin.H{"err": customErr.Err.Error(), "message": customErr.Error()})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": meeting})
}
