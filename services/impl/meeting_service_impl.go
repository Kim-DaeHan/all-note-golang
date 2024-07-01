package impl

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/models"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/Kim-DaeHan/all-note-golang/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MeetingServiceImpl struct {
	collection *mongo.Collection
}

func NewMeetingServiceImpl(collection *mongo.Collection) services.MeetingService {
	return &MeetingServiceImpl{collection}
}

func (ms *MeetingServiceImpl) GetMeeting(id string) (*models.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	meetingId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("Meeting", err)
	}

	var meeting *models.Meeting

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: meetingId}}}}

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "created_by"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "created_by_info"},
	}}}

	// Unwind participants 배열
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$participants"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}

	lookupParticipantsStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "participants.participant"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "participants.participant_info"},
		}},
	}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$_id"},
			{Key: "title", Value: bson.D{{Key: "$first", Value: "$title"}}},
			{Key: "description", Value: bson.D{{Key: "$first", Value: "$description"}}},
			{Key: "participants", Value: bson.D{{Key: "$push", Value: "$participants"}}},
			{Key: "start_dt", Value: bson.D{{Key: "$first", Value: "$start_dt"}}},
			{Key: "end_dt", Value: bson.D{{Key: "$first", Value: "$end_dt"}}},
			{Key: "location", Value: bson.D{{Key: "$first", Value: "$location"}}},
			{Key: "created_by", Value: bson.D{{Key: "$first", Value: "$created_by"}}},
			{Key: "created_by_info", Value: bson.D{{Key: "$first", Value: "$created_by_info"}}},
			{Key: "created_at", Value: bson.D{{Key: "$first", Value: "$created_at"}}},
			{Key: "updated_at", Value: bson.D{{Key: "$first", Value: "$updated_at"}}},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, lookupUserStage, unwindStage, lookupParticipantsStage, groupStage}

	result, err := ms.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer result.Close(ctx)

	if result.Next(ctx) {
		if err := result.Decode(&meeting); err != nil {
			return nil, &errors.CustomError{
				Message:    "결과 디코딩 오류",
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return meeting, nil
}

func (ms *MeetingServiceImpl) CreateMeeting(dto *dto.MeetingCreateDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("dto: %+v", dto)

	meeting := models.Meeting{
		ID:          primitive.NewObjectID(),
		Title:       dto.Title,
		Description: dto.Description,
		StartDt:     dto.StartDt,
		EndDt:       dto.EndDt,
		Location:    dto.Location,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	var err error

	if meeting.User, err = utils.ConvertToObjectId(dto.CreatedBy); err != nil {
		return utils.ConvertError("User", err)
	}

	// Participants 필드를 ObjectId로 변환하여 한 번에 할당
	var participantIDs []primitive.ObjectID
	if participantIDs, err = utils.ConvertStringIDsToObjectIDs(dto.Participants); err != nil {
		return utils.ConvertError("User", err)
	}

	meeting.Participants = make([]models.Participant, len(participantIDs))
	for i, objID := range participantIDs {
		meeting.Participants[i] = models.Participant{User: objID}
	}

	fmt.Printf("meeting: %+v", meeting)

	_, err = ms.collection.InsertOne(ctx, meeting)

	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}
