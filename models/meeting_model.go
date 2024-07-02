package models

import (
	"encoding/json"
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
	UserInfo     []meetingUser      `bson:"created_by_info,omitempty" json:"created_by_info,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
} //@name Meeting

type Participant struct {
	User     primitive.ObjectID `bson:"participant,omitempty" json:"participant,omitempty"`
	UserInfo []meetingUser      `bson:"participant_info,omitempty" json:"participant_info,omitempty"`
}

// meetingUser info
// @Description meetingUser information
type meetingUser struct {
	Email    string `bson:"email" json:"email"`
	UserName string `bson:"user_name" json:"user_name"`
	Position string `json:"position,omitempty"`
	Photo    string `json:"photo"`
} //@name meetingUser

// MarshalJSON() 함수는 Go의 encoding/json 패키지에서 제공하는 인터페이스 메서드입니다. 이 메서드를 구조체에 구현하면, 해당 구조체가 JSON으로 마샬링될 때 자동으로 호출
// 커스텀 JSON 마샬러 for Meeting
func (m Meeting) MarshalJSON() ([]byte, error) {
	type Alias Meeting
	aux := &struct {
		EndDt        *time.Time    `json:"end_dt,omitempty"`
		Participants []Participant `json:"participants,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&m),
	}

	// EndDt 필드가 기본값이면 nil로 설정
	if m.EndDt.IsZero() {
		aux.EndDt = nil
	} else {
		aux.EndDt = &m.EndDt
	}

	// 비어 있는 Participants 요소 제거
	for _, participant := range m.Participants {
		if !participant.isEmpty() {
			aux.Participants = append(aux.Participants, participant)
		}
	}

	return json.Marshal(aux)
}

// Participant 구조체가 비어 있는지 확인하는 메서드
func (p Participant) isEmpty() bool {
	return p.User.IsZero() && len(p.UserInfo) == 0
}

// 커스텀 JSON 마샬러 for Participant
func (p Participant) MarshalJSON() ([]byte, error) {
	type Alias Participant
	aux := &struct {
		User *string `json:"participant,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&p),
	}

	// User 필드가 기본값이면 nil로 설정
	if p.User.IsZero() {
		aux.User = nil
	} else {
		userStr := p.User.Hex()
		aux.User = &userStr
	}
	return json.Marshal(aux)
}
