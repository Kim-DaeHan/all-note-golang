package dto

import "time"

// MeetingCreateDTO info
// @Description Meeting information create dto
type MeetingCreateDTO struct {
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	Participants []string  `json:"participants,omitempty"`
	StartDt      time.Time `json:"start_dt"`
	EndDt        time.Time `json:"end_dt,omitempty"`
	Location     string    `json:"location,omitempty"`
	CreatedBy    string    `json:"created_by"`
} //@name MeetingCreateDTO
