package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type ProjectService interface {
	GetAllProject() ([]models.Project, error)
	CreateProject(dto *dto.ProjectCreateDTO) error
	UpdateProject(id string, dto *dto.ProjectUpdateDTO) (*models.Project, error)
	DeleteProject(id string) error
}
