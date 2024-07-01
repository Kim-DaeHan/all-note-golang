package dto

import "time"

// MeetingCreateDTO info
// @Description Meeting information create dto
type MeetingCreateDTO struct {
	Title        string    `json:"task"`
	Description  string    `json:"status"`
	Participants string    `json:"project,omitempty"`
	StartDt      time.Time `json:"start_dt"`
	EndDt        time.Time `json:"end_dt"`
	Location     string    `json:"department,omitempty"`
	User         string    `json:"user"`
} //@name MeetingCreateDTO
