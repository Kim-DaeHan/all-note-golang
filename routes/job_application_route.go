package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

type JobApplicationRoutes struct {
	jobApplicationHandler handlers.JobApplicationHandler
}

func NewJobApplicationRoutes(jobApplication handlers.JobApplicationHandler) JobApplicationRoutes {
	return JobApplicationRoutes{jobApplication}
}

func (jr *JobApplicationRoutes) SetJobApplicationRoutes(router *gin.RouterGroup) {
	jobApplications := router.Group("/jobApplications")

	jobApplications.GET("/", jr.jobApplicationHandler.GetAllJobApplication)
	jobApplications.GET("/:id", jr.jobApplicationHandler.GetJobApplication)
	jobApplications.GET("/manager/:id", jr.jobApplicationHandler.GetJobApplicationByManager)
	jobApplications.POST("/", jr.jobApplicationHandler.CreateJobApplication)
	jobApplications.PATCH("/:id", jr.jobApplicationHandler.UpdateJobApplication)
	jobApplications.DELETE("/:id", jr.jobApplicationHandler.DeleteJobApplication)

}
