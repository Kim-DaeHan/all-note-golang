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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MeetingServiceImpl struct {
	collection *mongo.Collection
}

func NewMeetingServiceImpl(collection *mongo.Collection) services.MeetingService {
	return &MeetingServiceImpl{collection}
}

func (ms *MeetingServiceImpl) GetAllMeeting() ([]models.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var meetings []models.Meeting

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

	pipeline := mongo.Pipeline{lookupUserStage, unwindStage, lookupParticipantsStage, groupStage}

	results, err := ms.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &meetings); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return meetings, nil
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

func (ms *MeetingServiceImpl) GetMeetingByUser(id string) ([]models.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("User", err)
	}

	var meetings []models.Meeting

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "created_by", Value: userId}}}}

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

	results, err := ms.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &meetings); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return meetings, nil
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

func (ms *MeetingServiceImpl) UpdateMeeting(id string, dto *dto.MeetingUpdateDTO) (*models.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	meetingId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("Meeting", err)
	}

	meeting := bson.M{
		"updated_at": time.Now(),
	}

	if dto.Title != "" {
		meeting["title"] = dto.Title
	}

	if dto.Description != "" {
		meeting["description"] = dto.Description
	}

	if dto.Location != "" {
		meeting["location"] = dto.Location
	}

	if !dto.StartDt.IsZero() {
		meeting["start_dt"] = dto.StartDt
	}

	if !dto.EndDt.IsZero() {
		meeting["end_dt"] = dto.EndDt
	}

	if len(dto.Participants) > 0 {
		participantIDs, err := utils.ConvertStringIDsToObjectIDs(dto.Participants)
		if err != nil {
			return nil, utils.ConvertError("Participant", err)
		}

		participants := make([]models.Participant, len(participantIDs))
		for i, objID := range participantIDs {
			participants[i] = models.Participant{User: objID}
		}

		meeting["participants"] = participants
	}

	filter := bson.M{"_id": meetingId}
	update := bson.M{"$set": meeting}

	fmt.Printf("meeting: %+v", meeting)

	result := ms.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, &errors.CustomError{
				Message:    "Meeting을 찾을 수 없음",
				StatusCode: http.StatusNotFound,
				Err:        result.Err(),
			}
		}
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        result.Err(),
		}
	}

	var updatedMeeting *models.Meeting
	if err := result.Decode(&updatedMeeting); err != nil {
		return nil, &errors.CustomError{
			Message:    "결과 디코딩 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return updatedMeeting, nil
}

func (ms *MeetingServiceImpl) DeleteMeeting(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	meetingId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return utils.ConvertError("Meeting", err)
	}

	filter := bson.M{"_id": meetingId}

	result, err := ms.collection.DeleteOne(ctx, filter)
	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if result.DeletedCount == 0 {
		return &errors.CustomError{
			Message:    "Meeting을 찾을 수 없음",
			StatusCode: http.StatusNotFound,
			Err:        mongo.ErrNoDocuments,
		}
	}

	return nil
}
