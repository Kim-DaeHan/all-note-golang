package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	noteService services.NoteService
}

func NewNoteController(noteService services.NoteService) NoteHandler {
	return NoteHandler{noteService}
}

// GetAllNote godoc
// @Summary 전체 노트 조회
// @Description 전체 노트 조회
// @name GetAllNote
// @Accept  json
// @Produce  json
// @Router /notes [get]
// @Success 200 {object} Note
// @Failure 500
func (nh *NoteHandler) GetAllNote(ctx *gin.Context) {
	notes, err := nh.noteService.GetAllNote()

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

	ctx.IndentedJSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": notes})
}

// GetNote godoc
// @Summary 노트 조회
// @Description 노트 조회
// @name GetNote
// @Accept  json
// @Produce  json
// @Param noteId path string true "Note ID"
// @Router /notes/{noteId} [get]
// @Success 200 {object} Note
// @Failure 500
func (nh *NoteHandler) GetNote(ctx *gin.Context) {
	id := ctx.Param("id")

	note, err := nh.noteService.GetNote(id)

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

	ctx.IndentedJSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": note})
}

// GetNoteByUser godoc
// @Summary 노트 조회(유저)
// @Description 노트 조회(유저)
// @name GetNoteByUser
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID"
// @Router /notes/{userId}/user [get]
// @Success 200 {object} Note
// @Failure 500
func (nh *NoteHandler) GetNoteByUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	notes, err := nh.noteService.GetNoteByUser(userId)

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

	ctx.IndentedJSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": notes})
}

// CreateNote godoc
// @Summary 노트 생성
// @Description 노트 생성
// @name CreateNote
// @Accept  json
// @Produce  json
// @Param note body dto.NoteCreateDTO true "노트 정보"
// @Router /notes [post]
// @Success 200 {object} Note
// @Failure 500
func (nh *NoteHandler) CreateNote(ctx *gin.Context) {
	var dto dto.NoteCreateDTO

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

	result, err := nh.noteService.CreateNote(dto)

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

	ctx.IndentedJSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": result})
}
