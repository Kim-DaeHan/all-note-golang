package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

type JobApplicationHandler struct {
	jobApplicationService services.JobApplicationService
}

func NewJobApplicationHandler(jobApplicationService services.JobApplicationService) JobApplicationHandler {
	return JobApplicationHandler{jobApplicationService}
}

// GetAllJobApplication godoc
// @Tags JobApplication
// @Summary 전체 JobApplication 조회
// @Description 전체 JobApplication 조회
// @ID GetAllJobApplication
// @Accept  json
// @Produce  json
// @Router /jobApplications [get]
// @Success 200 {object} dto.APIResponse[[]JobApplication]
// @Failure 500
func (jh *JobApplicationHandler) GetAllJobApplication(ctx *gin.Context) {
	jobApplications, err := jh.jobApplicationService.GetAllJobApplication()

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": jobApplications})
}

// GetJobApplication godoc
// @Tags JobApplication
// @Summary JobApplication 조회
// @Description JobApplication 조회
// @ID GetJobApplication
// @Accept  json
// @Produce  json
// @Param jobApplicationId path string true "JobApplication ID"
// @Router /jobApplications/{jobApplicationId} [get]
// @Success 200 {object} dto.APIResponse[JobApplication]
// @Failure 500
func (jh *JobApplicationHandler) GetJobApplication(ctx *gin.Context) {
	jobApplicationId := ctx.Param("id")

	jobApplication, err := jh.jobApplicationService.GetJobApplication(jobApplicationId)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": jobApplication})
}

// GetJobApplicationByManager godoc
// @Tags JobApplication
// @Summary JobApplication 조회(담장자)
// @Description JobApplication 조회(담당자)
// @ID GetJobApplicationByManager
// @Accept  json
// @Produce  json
// @Param managerId path string true "Manager User ID"
// @Router /jobApplications/manager/{managerId} [get]
// @Success 200 {object} dto.APIResponse[[]JobApplication]
// @Failure 500
func (jh *JobApplicationHandler) GetJobApplicationByManager(ctx *gin.Context) {
	managerId := ctx.Param("id")

	jobApplications, err := jh.jobApplicationService.GetJobApplicationByManager(managerId)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": jobApplications})
}

// CreateJobApplication godoc
// @Tags JobApplication
// @Summary JobApplication 생성
// @Description JobApplication 생성
// @ID CreateJobApplication
// @Accept  json
// @Produce  json
// @Param jobApplication body dto.JobApplicationCreateDTO true "JobApplication 정보"
// @Router /jobApplications [post]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (jh *JobApplicationHandler) CreateJobApplication(ctx *gin.Context) {
	var dto dto.JobApplicationCreateDTO

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

	err := jh.jobApplicationService.CreateJobApplication(&dto)

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

// UpdateJobApplication godoc
// @Tags JobApplication
// @Summary JobApplication 수정
// @Description JobApplication 수정
// @ID UpdateJobApplication
// @Accept  json
// @Produce  json
// @Param jobApplicationId path string true "JobApplication ID"
// @Param jobApplication body dto.JobApplicationUpdateDTO true "JobApplication 정보"
// @Router /jobApplications/{jobApplicationId} [patch]
// @Success 200 {object} dto.APIResponse[JobApplication]
// @Failure 500
func (jh *JobApplicationHandler) UpdateJobApplication(ctx *gin.Context) {
	var dto dto.JobApplicationUpdateDTO
	jobApplicationId := ctx.Param("id")

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

	jobApplication, err := jh.jobApplicationService.UpdateJobApplication(jobApplicationId, &dto)

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

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": jobApplication})
}

// DeleteJobApplication godoc
// @Tags JobApplication
// @Summary JobApplication 삭제
// @Description JobApplication 삭제
// @ID DeleteJobApplication
// @Accept  json
// @Produce  json
// @Param jobApplicationId path string true "JobApplication ID"
// @Router /jobApplications/{jobApplicationId} [delete]
// @Success 200 {object} dto.APIResponseWithoutData
// @Failure 500
func (jh *JobApplicationHandler) DeleteJobApplication(ctx *gin.Context) {
	jobApplicationId := ctx.Param("id")

	err := jh.jobApplicationService.DeleteJobApplication(jobApplicationId)

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
