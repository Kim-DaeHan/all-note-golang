package dto

import "time"

// JobApplicationUpdateDTO info
// @Description JobApplication information update dto
type JobApplicationUpdateDTO struct {
	ApplicantName string    `json:"applicant_name"`
	User          string    `json:"manager"`
	Department    string    `json:"department,omitempty"`
	Position      string    `json:"position"`
	Task          string    `json:"task"`
	Stage         string    `json:"stage"`
	Location      string    `json:"location,omitempty"`
	Status        string    `json:"status"`
	StartDt       time.Time `json:"start_dt"`
	EndDt         time.Time `json:"end_dt,omitempty"`
} //@name JobApplicationUpdateDTO
