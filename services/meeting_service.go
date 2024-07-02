package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type MeetingService interface {
	GetAllMeeting() ([]models.Meeting, error)
	GetMeeting(id string) (*models.Meeting, error)
	GetMeetingByUser(userId string) ([]models.Meeting, error)
	CreateMeeting(dto *dto.MeetingCreateDTO) error
	UpdateMeeting(id string, dto *dto.MeetingUpdateDTO) (*models.Meeting, error)
	DeleteMeeting(id string) error
}
