package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type JobApplicationService interface {
	GetAllJobApplication() ([]models.JobApplication, error)
	GetJobApplication(id string) (*models.JobApplication, error)
	GetJobApplicationByManager(userId string) ([]models.JobApplication, error)
	CreateJobApplication(dto *dto.JobApplicationCreateDTO) error
	UpdateJobApplication(id string, dto *dto.JobApplicationUpdateDTO) (*models.JobApplication, error)
	DeleteJobApplication(id string) error
}
