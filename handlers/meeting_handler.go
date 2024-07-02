package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/dto"
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

// GetAllMeeting godoc
// @Tags Meeting
// @Summary 전체 Meeting 조회
// @Description 전체 Meeting 조회
// @ID GetAllMeeting
// @Accept  json
// @Produce  json
// @Router /meetings [get]
// @Success 200 {object} dto.APIResponse[[]Meeting]
// @Failure 500
func (mh *MeetingHandler) GetAllMeeting(ctx *gin.Context) {
	meetings, err := mh.meetingService.GetAllMeeting()

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": meetings})
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

// GetMeetingByUser godoc
// @Tags Meeting
// @Summary Meeting 조회(유저)
// @Description Meeting 조회(유저)
// @ID GetMeetingByUser
// @Accept  json
// @Produce  json
// @Param userId path string true "Created By ID"
// @Router /meetings/created-by/{userId} [get]
// @Success 200 {object} dto.APIResponse[[]Meeting]
// @Failure 500
func (mh *MeetingHandler) GetMeetingByUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	meetings, err := mh.meetingService.GetMeetingByUser(userId)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": meetings})
}

// CreateMeeting godoc
// @Tags Meeting
// @Summary Meeting 생성
// @Description Meeting 생성
// @ID CreateMeeting
// @Accept  json
// @Produce  json
// @Param meeting body dto.MeetingCreateDTO true "Meeting 정보"
// @Router /meetings [post]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (mh *MeetingHandler) CreateMeeting(ctx *gin.Context) {
	var dto dto.MeetingCreateDTO

	//validate the request body
	if err := ctx.BindJSON(&dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&dto); validationErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": validationErr.Error()})
		return
	}

	err := mh.meetingService.CreateMeeting(&dto)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully"})
}

// UpdateMeeting godoc
// @Tags Meeting
// @Summary Meeting 수정
// @Description Meeting 수정
// @ID UpdateMeeting
// @Accept  json
// @Produce  json
// @Param meetingId path string true "Meeting ID"
// @Param meeting body dto.MeetingUpdateDTO true "Meeting 정보"
// @Router /meetings/{meetingId} [patch]
// @Success 200 {object} dto.APIResponse[Meeting]
// @Failure 500
func (mh *MeetingHandler) UpdateMeeting(ctx *gin.Context) {
	var dto dto.MeetingUpdateDTO
	meetingId := ctx.Param("id")

	//validate the request body
	if err := ctx.BindJSON(&dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&dto); validationErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": validationErr.Error()})
		return
	}

	meeting, err := mh.meetingService.UpdateMeeting(meetingId, &dto)

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

// DeleteMeeting godoc
// @Tags Meeting
// @Summary Meeting 삭제
// @Description Meeting 삭제
// @ID DeleteMeeting
// @Accept  json
// @Produce  json
// @Param meetingId path string true "Meeting ID"
// @Router /meetings/{meetingId} [delete]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (mh *MeetingHandler) DeleteMeeting(ctx *gin.Context) {
	meetingId := ctx.Param("id")

	err := mh.meetingService.DeleteMeeting(meetingId)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully"})
}
