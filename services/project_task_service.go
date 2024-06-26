package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type ProjectTaskService interface {
	GetProjectTask(id string) (*models.ProjectTask, error)
	GetProjectTaskByProject(userId string) ([]models.ProjectTask, error)
	CreateProjectTask(dto *dto.ProjectTaskCreateDTO) error
	UpdateProjectTask(id string, dto *dto.ProjectTaskUpdateDTO) (*models.ProjectTask, error)
	DeleteProjectTask(id string) error
}
