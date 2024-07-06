package models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JobApplication info
// @Description JobApplication information
type JobApplication struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	ApplicantName  string             `bson:"applicant_name" json:"applicant_name"`
	User           primitive.ObjectID `bson:"manager" json:"manager"`
	UserInfo       []jobUser          `bson:"manager_info,omitempty" json:"manager_info,omitempty"`
	Department     primitive.ObjectID `bson:"department,omitempty" json:"department,omitempty"`
	DepartmentInfo []Department       `bson:"department_info,omitempty" json:"department_info,omitempty"`
	Position       string             `bson:"position" json:"position"`
	Task           string             `bson:"task" json:"task"`
	Stage          string             `bson:"stage" json:"stage"`
	Location       string             `bson:"location,omitempty" json:"location,omitempty"`
	StartDt        time.Time          `bson:"start_dt" json:"start_dt"`
	EndDt          time.Time          `bson:"end_dt,omitempty" json:"end_dt,omitempty"`
	Status         string             `bson:"status" json:"status"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
} //@name JobApplication

// meetingUser info
// @Description meetingUser information
type jobUser struct {
	Email    string `bson:"email" json:"email"`
	UserName string `bson:"user_name" json:"user_name"`
	Position string `json:"position,omitempty"`
	Photo    string `json:"photo"`
} //@name meetingUser

func (j *JobApplication) MarshalJSON() ([]byte, error) {
	type Alias JobApplication
	aux := &struct {
		*Alias
		EndDt *time.Time `json:"end_dt,omitempty"`
	}{
		Alias: (*Alias)(j),
	}
	if j.EndDt.IsZero() {
		aux.EndDt = nil
	} else {
		aux.EndDt = &j.EndDt
	}

	return json.Marshal(aux)
}
