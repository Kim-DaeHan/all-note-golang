package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService services.ProjectService
}

func NewProjectHandler(projectService services.ProjectService) ProjectHandler {
	return ProjectHandler{projectService}
}

// GetAllProject godoc
// @Tags Project
// @Summary 전체 Project 조회
// @Description 전체 Project 조회
// @ID GetAllProject
// @Accept  json
// @Produce  json
// @Router /projects [get]
// @Success 200 {object} dto.APIResponse[[]Project]
// @Failure 500
func (ph *ProjectHandler) GetAllProject(ctx *gin.Context) {
	projects, err := ph.projectService.GetAllProject()

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": projects})
}

// CreateProject godoc
// @Tags Project
// @Summary Project 생성
// @Description Project 생성
// @ID CreateProject
// @Accept  json
// @Produce  json
// @Param project body dto.ProjectCreateDTO true "Project 정보"
// @Router /projects [post]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (ph *ProjectHandler) CreateProject(ctx *gin.Context) {
	var dto dto.ProjectCreateDTO

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

	err := ph.projectService.CreateProject(&dto)

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

// UpdateProject godoc
// @Tags Project
// @Summary Project 수정
// @Description Project 수정
// @ID UpdateProject
// @Accept  json
// @Produce  json
// @Param projectId path string true "Project ID"
// @Param project body dto.ProjectUpdateDTO true "Project 정보"
// @Router /projects/{projectId} [patch]
// @Success 200 {object} dto.APIResponse[Project]
// @Failure 500
func (ph *ProjectHandler) UpdateProject(ctx *gin.Context) {
	var dto dto.ProjectUpdateDTO
	projectId := ctx.Param("id")

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

	project, err := ph.projectService.UpdateProject(projectId, &dto)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": project})
}

// DeleteProject godoc
// @Tags Project
// @Summary Project 삭제
// @Description Project 삭제
// @ID DeleteProject
// @Accept  json
// @Produce  json
// @Param projectId path string true "Project ID"
// @Router /projects/{projectId} [delete]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (ph *ProjectHandler) DeleteProject(ctx *gin.Context) {
	projectId := ctx.Param("id")

	err := ph.projectService.DeleteProject(projectId)

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
