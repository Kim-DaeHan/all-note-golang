package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserHandler {
	return UserHandler{userService}
}

// GetAllUser godoc
// @Tags User
// @Summary 전체 유저 조회
// @Description 전체 유저 조회
// @ID GetAllUser
// @Accept  json
// @Produce  json
// @Router /users [get]
// @Success 200 {object} dto.APIResponse[[]User]
// @Failure 500
func (uh *UserHandler) GetAllUser(ctx *gin.Context) {
	users, err := uh.userService.GetAllUser()

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": users})
}

// GetUser godoc
// @Tags User
// @Summary 유저 조회
// @Description 유저 조회
// @ID GetUser
// @Accept  json
// @Produce  json
// @Param userId path string true "유저 ID"
// @Router /users/{userId} [get]
// @Success 200 {object} dto.APIResponse[User]
// @Failure 500
func (uh *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	users, err := uh.userService.GetUser(id)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": users})
}

// CreateUser godoc
// @Tags User
// @Summary 유저 생성
// @Description 유저 생성
// @ID CreateUser
// @Accept  json
// @Produce  json
// @Param user body dto.UserCreateDTO true "유저 정보"
// @Router /users [post]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var dto dto.UserCreateDTO

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

	err := uh.userService.CreateUser(&dto)

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

// UpsertUser godoc
// @Tags User
// @Summary 유저 Upsert
// @Description 유저 생성 or 업데이트
// @ID UpsertUser
// @Accept  json
// @Produce  json
// @Param user body dto.UserUpdateDTO true "유저 정보"
// @Router /users/upsert [post]
// @Success 200 {object} dto.APIResponse[User]
// @Failure 500
func (uh *UserHandler) UpsertUser(ctx *gin.Context) {
	var dto dto.UserUpdateDTO

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

	result, err := uh.userService.UpsertUser(&dto)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": result})
}
