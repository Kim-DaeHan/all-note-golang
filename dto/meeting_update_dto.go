package dto

import "time"

// MeetingUpdateDTO info
// @Description Meeting information update dto
type MeetingUpdateDTO struct {
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	Participants []string  `json:"participants,omitempty"`
	StartDt      time.Time `json:"start_dt,omitempty"`
	EndDt        time.Time `json:"end_dt,omitempty"`
	Location     string    `json:"location,omitempty"`
} //@name MeetingUpdateDTO
