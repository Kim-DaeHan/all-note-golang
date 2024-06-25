package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) TodoHandler {
	return TodoHandler{todoService}
}

// GetAllTodo godoc
// @Tags Todo
// @Summary 전체 Todo 조회
// @Description 전체 Todo 조회
// @ID GetAllTodo
// @Accept  json
// @Produce  json
// @Router /todos [get]
// @Success 200 {object} dto.APIResponse[[]Todo]
// @Failure 500
func (th *TodoHandler) GetAllTodo(ctx *gin.Context) {
	todos, err := th.todoService.GetAllTodo()

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": todos})
}

// GetTodo godoc
// @Tags Todo
// @Summary Todo 조회
// @Description Todo 조회
// @ID GetTodo
// @Accept  json
// @Produce  json
// @Param todoId path string true "Todo ID"
// @Router /todos/{todoId} [get]
// @Success 200 {object} dto.APIResponse[Todo]
// @Failure 500
func (th *TodoHandler) GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := th.todoService.GetTodo(id)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": todo})
}

// GetTodoByUser godoc
// @Tags Todo
// @Summary Todo 조회(유저)
// @Description Todo 조회(유저)
// @ID GetTodoByUser
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID"
// @Router /todos/{userId}/user [get]
// @Success 200 {object} dto.APIResponse[[]Todo]
// @Failure 500
func (th *TodoHandler) GetTodoByUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	todos, err := th.todoService.GetTodoByUser(userId)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": todos})
}

// CreateTodo godoc
// @Tags Todo
// @Summary Todo 생성
// @Description Todo 생성
// @ID CreateTodo
// @Accept  json
// @Produce  json
// @Param todo body dto.TodoCreateDTO true "Todo 정보"
// @Router /todos [post]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (th *TodoHandler) CreateTodo(ctx *gin.Context) {
	var dto dto.TodoCreateDTO

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

	err := th.todoService.CreateTodo(&dto)

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

// UpdateTodo godoc
// @Tags Todo
// @Summary Todo 수정
// @Description Todo 수정
// @ID UpdateTodo
// @Accept  json
// @Produce  json
// @Param todoId path string true "Todo ID"
// @Param todo body dto.TodoUpdateDTO true "Todo 정보"
// @Router /todos/{todoId} [patch]
// @Success 200 {object} dto.APIResponse[Todo]
// @Failure 500
func (th *TodoHandler) UpdateTodo(ctx *gin.Context) {
	var dto dto.TodoUpdateDTO
	todoId := ctx.Param("id")

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

	todo, err := th.todoService.UpdateTodo(todoId, &dto)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": todo})
}

// DeleteTodo godoc
// @Tags Todo
// @Summary Todo 삭제
// @Description Todo 삭제
// @ID DeleteTodo
// @Accept  json
// @Produce  json
// @Param todoId path string true "Todo ID"
// @Router /todos/{todoId} [delete]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (th *TodoHandler) DeleteTodo(ctx *gin.Context) {
	todoId := ctx.Param("id")

	err := th.todoService.DeleteTodo(todoId)

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
