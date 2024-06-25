package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

type ProjectRoutes struct {
	projectHandler handlers.ProjectHandler
}

func NewProjectRoutes(projectHandler handlers.ProjectHandler) ProjectRoutes {
	return ProjectRoutes{projectHandler}
}

func (pr *ProjectRoutes) SetProjectRoutes(router *gin.RouterGroup) {
	projects := router.Group("/projects")

	projects.GET("/", pr.projectHandler.GetAllProject)
	projects.POST("/", pr.projectHandler.CreateProject)
	projects.PATCH("/:id", pr.projectHandler.UpdateProject)
	projects.DELETE("/:id", pr.projectHandler.DeleteProject)

}
