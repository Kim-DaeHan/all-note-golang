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

type NoteServiceImpl struct {
	collection *mongo.Collection
}

func NewNoteServiceImpl(collection *mongo.Collection) services.NoteService {
	return &NoteServiceImpl{collection}
}

func (ns *NoteServiceImpl) GetAllNote() ([]models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var notes []models.Note

	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "author"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "author_info"},
	}}}

	pipeline := mongo.Pipeline{lookupStage}

	// query := bson.M{}
	// results, err := us.collection.Find(ctx, query)
	results, err := ns.collection.Aggregate(ctx, pipeline)
	// err := us.collection.FindOne(ctx, query).Decode(&users)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)

	if err = results.All(ctx, &notes); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return notes, nil
}

func (ns *NoteServiceImpl) GetNote(id string) (*models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	noteId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("Todo", err)
	}

	var note *models.Note

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: noteId}}}}

	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "author"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "author_info"},
	}}}

	pipeline := mongo.Pipeline{matchStage, lookupStage}

	result, err := ns.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer result.Close(ctx)

	if result.Next(ctx) {
		if err := result.Decode(&note); err != nil {
			return nil, &errors.CustomError{
				Message:    "결과 디코딩 오류",
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return note, nil
}

func (ns *NoteServiceImpl) GetNoteByUser(id string) ([]models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("User", err)
	}

	var notes []models.Note

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "author", Value: userId}}}}

	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "author"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "author_info"},
	}}}

	pipeline := mongo.Pipeline{matchStage, lookupStage}

	results, err := ns.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &notes); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return notes, nil
}

func (ns *NoteServiceImpl) CreateNote(dto *dto.NoteCreateDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("dto: %+v", dto)

	note := models.Note{
		ID:        primitive.NewObjectID(),
		Text:      dto.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var err error

	if note.Author, err = utils.ConvertToObjectId(dto.Author); err != nil {
		return utils.ConvertError("Note", err)
	}

	fmt.Printf("note: %+v", note)

	_, err = ns.collection.InsertOne(ctx, note)

	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}

func (ns *NoteServiceImpl) UpdateNote(id string, dto *dto.NoteUpdateDTO) (*models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	noteId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("Note", err)
	}

	note := bson.M{
		"text":       dto.Text,
		"updated_at": time.Now(),
	}

	filter := bson.M{"_id": noteId}
	update := bson.M{"$set": note}

	result := ns.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, &errors.CustomError{
				Message:    "노트를 찾을 수 없음",
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

	var updatedNote *models.Note
	if err := result.Decode(&updatedNote); err != nil {
		return nil, &errors.CustomError{
			Message:    "결과 디코딩 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return updatedNote, nil
}

func (ns *NoteServiceImpl) DeleteNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &errors.CustomError{
			Message:    "Note ObjectID 변환 오류",
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	filter := bson.M{"_id": objID}

	result, err := ns.collection.DeleteOne(ctx, filter)
	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if result.DeletedCount == 0 {
		return &errors.CustomError{
			Message:    "노트를 찾을 수 없음",
			StatusCode: http.StatusNotFound,
			Err:        mongo.ErrNoDocuments,
		}
	}

	return nil
}
