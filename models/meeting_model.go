package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Meeting info
// @Description Meeting information
type Meeting struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description,omitempty" json:"description,omitempty"`
	Participants []Participant      `bson:"participants,omitempty" json:"participants,omitempty"`
	StartDt      time.Time          `bson:"start_dt" json:"start_dt"`
	EndDt        time.Time          `bson:"end_dt,omitempty" json:"end_dt,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	User         primitive.ObjectID `bson:"created_by" json:"created_by"`
	UserInfo     []User             `bson:"created_by_info,omitempty" json:"created_by_info,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
} //@name Meeting

type Participant struct {
	User     primitive.ObjectID `bson:"participant" json:"participant"`
	UserInfo []User             `bson:"participant_info,omitempty" json:"participant_info,omitempty"`
}
