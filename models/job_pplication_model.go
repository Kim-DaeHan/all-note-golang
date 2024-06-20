package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JobApplication info
// @Description JobApplication information
type JobApplication struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Name           string             `bson:"name" json:"name"`
	User           primitive.ObjectID `bson:"manager" json:"manager"`
	UserInfo       []User             `bson:"manager_info,omitempty" json:"manager_info,omitempty"`
	Department     primitive.ObjectID `bson:"department,omitempty" json:"department,omitempty"`
	DepartmentInfo []Department       `bson:"department_info,omitempty" json:"department_info,omitempty"`
	Position       string             `bson:"position" json:"position"`
	Task           string             `bson:"task" json:"task"`
	Stage          string             `bson:"stage" json:"stage"`
	Location       string             `bson:"location" json:"location"`
	StartDt        time.Time          `bson:"start_dt" json:"start_dt"`
	EndDt          time.Time          `bson:"end_dt" json:"end_dt"`
	Status         string             `bson:"status" json:"status"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
} //@name JobApplication
