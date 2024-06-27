package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

type ProjectTaskRoutes struct {
	projectTaskHandler handlers.ProjectTaskHandler
}

func NewProjectTaskRoutes(projectTaskHandler handlers.ProjectTaskHandler) ProjectTaskRoutes {
	return ProjectTaskRoutes{projectTaskHandler}
}

func (ptr *ProjectTaskRoutes) SetProjectTaskRoutes(router *gin.RouterGroup) {
	tasks := router.Group("/project-tasks")

	tasks.GET("/:id", ptr.projectTaskHandler.GetProjectTask)
	tasks.GET("/project/:id", ptr.projectTaskHandler.GetProjectTaskByProject)
	tasks.POST("/", ptr.projectTaskHandler.CreateProjectTask)
	tasks.PATCH("/:id", ptr.projectTaskHandler.UpdateProjectTask)
	tasks.DELETE("/:id", ptr.projectTaskHandler.DeleteProjectTask)

}
