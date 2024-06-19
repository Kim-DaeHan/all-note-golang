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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoServiceImpl struct {
	collection *mongo.Collection
}

func NewTodoServiceImpl(collection *mongo.Collection) services.TodoService {
	return &TodoServiceImpl{collection}
}

func (ts *TodoServiceImpl) GetAllTodo() ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var todos []models.Todo

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "user"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "user_info"},
	}}}

	lookupProjectStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "projects"},
		{Key: "localField", Value: "project"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "project_info"},
	}}}

	lookupDepartmentStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "departments"},
		{Key: "localField", Value: "department"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "department_info"},
	}}}

	pipeline := mongo.Pipeline{lookupUserStage, lookupProjectStage, lookupDepartmentStage}

	results, err := ts.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &todos); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return todos, nil
}

func (ts *TodoServiceImpl) GetTodo(id string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &errors.CustomError{
			Message:    "잘못된 ID 형식",
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	var todos *models.Todo

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: objID}}}}

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "user"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "user_info"},
	}}}

	lookupProjectStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "projects"},
		{Key: "localField", Value: "project"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "project_info"},
	}}}

	lookupDepartmentStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "departments"},
		{Key: "localField", Value: "department"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "department_info"},
	}}}

	pipeline := mongo.Pipeline{matchStage, lookupUserStage, lookupProjectStage, lookupDepartmentStage}

	result, err := ts.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer result.Close(ctx)

	if result.Next(ctx) {
		if err := result.Decode(&todos); err != nil {
			return nil, &errors.CustomError{
				Message:    "결과 디코딩 오류",
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return todos, nil
}

func (ts *TodoServiceImpl) GetTodoByUser(userId string) ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, &errors.CustomError{
			Message:    "잘못된 ID 형식",
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	var todos []models.Todo

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user", Value: objID}}}}

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "user"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "user_info"},
	}}}

	lookupProjectStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "projects"},
		{Key: "localField", Value: "project"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "project_info"},
	}}}

	lookupDepartmentStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "departments"},
		{Key: "localField", Value: "department"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "department_info"},
	}}}

	pipeline := mongo.Pipeline{matchStage, lookupUserStage, lookupProjectStage, lookupDepartmentStage}

	results, err := ts.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &todos); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return todos, nil
}

func (ts *TodoServiceImpl) CreateTodo(dto *dto.TodoCreateDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userId primitive.ObjectID

	fmt.Printf("dto: %+v", dto)

	note := models.Todo{
		ID:        primitive.NewObjectID(),
		Task:      dto.Task,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if dto.Author != "" {
		var err error
		userId, err = primitive.ObjectIDFromHex(dto.Author)

		if err != nil {
			return &errors.CustomError{
				Message:    "User ObjectID 변환 오류",
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}

		note.Author = userId
	}

	fmt.Printf("user: %+v", note)

	_, err := ns.collection.InsertOne(ctx, note)

	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}

func (ts *TodoServiceImpl) UpdateTodo(id string, dto *dto.TodoUpdateDTO) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &errors.CustomError{
			Message:    "Note ObjectID 변환 오류",
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	note := bson.M{
		"text":       dto.Text,
		"updated_at": time.Now(),
	}

	filter := bson.M{"_id": objID}
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

func (ts *TodoServiceImpl) DeleteTodo(id string) error {
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
