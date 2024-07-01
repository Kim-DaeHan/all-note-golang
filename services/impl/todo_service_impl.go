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

	todoId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("Todo", err)
	}

	var todo *models.Todo

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: todoId}}}}

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
		if err := result.Decode(&todo); err != nil {
			return nil, &errors.CustomError{
				Message:    "결과 디코딩 오류",
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return todo, nil
}

func (ts *TodoServiceImpl) GetTodoByUser(id string) ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("User", err)
	}

	var todos []models.Todo

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user", Value: userId}}}}

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

	fmt.Printf("dto: %+v", dto)

	todo := models.Todo{
		ID:        primitive.NewObjectID(),
		Task:      dto.Task,
		Status:    dto.Status,
		StartDt:   dto.StartDt,
		EndDt:     dto.EndDt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var err error

	if todo.Project, err = utils.ConvertToObjectId(dto.Project); err != nil {
		return utils.ConvertError("Project", err)
	}

	if todo.User, err = utils.ConvertToObjectId(dto.User); err != nil {
		return utils.ConvertError("User", err)
	}

	if todo.Department, err = utils.ConvertToObjectId(dto.Department); err != nil {
		return utils.ConvertError("Department", err)
	}

	fmt.Printf("todo: %+v", todo)

	_, err = ts.collection.InsertOne(ctx, todo)

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

	todoId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("Todo", err)
	}

	todo := bson.M{
		"updated_at": time.Now(),
	}

	if dto.Task != "" {
		todo["task"] = dto.Task
	}

	if dto.Status != "" {
		todo["status"] = dto.Status
	}

	if !dto.StartDt.IsZero() {
		todo["start_dt"] = dto.StartDt
	}

	if !dto.EndDt.IsZero() {
		todo["end_dt"] = dto.EndDt
	}

	if dto.Department != "" {
		if todo["department"], err = utils.ConvertToObjectId(dto.Department); err != nil {
			return nil, utils.ConvertError("Department", err)
		}
	}

	if dto.Project != "" {
		if todo["project"], err = utils.ConvertToObjectId(dto.Project); err != nil {
			return nil, utils.ConvertError("Project", err)
		}
	}

	filter := bson.M{"_id": todoId}
	update := bson.M{"$set": todo}

	fmt.Printf("todo: %+v", todo)

	result := ts.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, &errors.CustomError{
				Message:    "TODO를 찾을 수 없음",
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

	var updatedTodo *models.Todo
	if err := result.Decode(&updatedTodo); err != nil {
		return nil, &errors.CustomError{
			Message:    "결과 디코딩 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return updatedTodo, nil
}

func (ts *TodoServiceImpl) DeleteTodo(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todoId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return utils.ConvertError("Todo", err)
	}

	filter := bson.M{"_id": todoId}

	result, err := ts.collection.DeleteOne(ctx, filter)
	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if result.DeletedCount == 0 {
		return &errors.CustomError{
			Message:    "TODO를 찾을 수 없음",
			StatusCode: http.StatusNotFound,
			Err:        mongo.ErrNoDocuments,
		}
	}

	return nil
}
