package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

type ProjectTaskHandler struct {
	projectTaskService services.ProjectTaskService
}

func NewProjectTaskHandler(projectTaskService services.ProjectTaskService) ProjectTaskHandler {
	return ProjectTaskHandler{projectTaskService}
}

// GetProjectTask godoc
// @Tags ProjectTask
// @Summary ProjectTask 조회
// @Description ProjectTask 조회
// @ID GetProjectTask
// @Accept  json
// @Produce  json
// @Param taskId path string true "Project Task ID"
// @Router /project-tasks/{taskId} [get]
// @Success 200 {object} dto.APIResponse[ProjectTask]
// @Failure 500
func (pth *ProjectTaskHandler) GetProjectTask(ctx *gin.Context) {
	taskId := ctx.Param("id")

	task, err := pth.projectTaskService.GetProjectTask(taskId)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": task})
}

// GetProjectTaskByProject godoc
// @Tags ProjectTask
// @Summary ProjectTask 조회(프로젝트)
// @Description ProjectTask 조회(프로젝트)
// @ID GetProjectTaskByProject
// @Accept  json
// @Produce  json
// @Param projectId path string true "Project ID"
// @Router /project-tasks/project/{projectId} [get]
// @Success 200 {object} dto.APIResponse[[]ProjectTask]
// @Failure 500
func (pth *ProjectTaskHandler) GetProjectTaskByProject(ctx *gin.Context) {
	projectId := ctx.Param("id")

	tasks, err := pth.projectTaskService.GetProjectTaskByProject(projectId)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": tasks})
}

// CreateProjectTask godoc
// @Tags ProjectTask
// @Summary ProjectTask 생성
// @Description ProjectTask 생성
// @ID CreateProjectTask
// @Accept  json
// @Produce  json
// @Param projectTask body dto.ProjectTaskCreateDTO true "ProjectTask 정보"
// @Router /project-tasks [post]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (pth *ProjectTaskHandler) CreateProjectTask(ctx *gin.Context) {
	var dto dto.ProjectTaskCreateDTO

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

	err := pth.projectTaskService.CreateProjectTask(&dto)

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

// UpdateProjectTask godoc
// @Tags ProjectTask
// @Summary ProjectTask 수정
// @Description ProjectTask 수정
// @ID UpdateProjectTask
// @Accept  json
// @Produce  json
// @Param taskId path string true "Project Task ID"
// @Param projectTask body dto.ProjectTaskUpdateDTO true "ProjectTask 정보"
// @Router /project-tasks/{taskId} [patch]
// @Success 200 {object} dto.APIResponse[ProjectTask]
// @Failure 500
func (pth *ProjectTaskHandler) UpdateProjectTask(ctx *gin.Context) {
	var dto dto.ProjectTaskUpdateDTO
	taskId := ctx.Param("id")

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

	task, err := pth.projectTaskService.UpdateProjectTask(taskId, &dto)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": task})
}

// DeleteProjectTask godoc
// @Tags ProjectTask
// @Summary ProjectTask 삭제
// @Description ProjectTask 삭제
// @ID DeleteProjectTask
// @Accept  json
// @Produce  json
// @Param taskId path string true "ProjectTask ID"
// @Router /project-tasks/{taskId} [delete]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (pth *ProjectTaskHandler) DeleteProjectTask(ctx *gin.Context) {
	taskId := ctx.Param("id")

	err := pth.projectTaskService.DeleteProjectTask(taskId)

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
