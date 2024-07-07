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

type ProjectTaskServiceImpl struct {
	collection *mongo.Collection
}

func NewProjectTaskServiceImpl(collection *mongo.Collection) services.ProjectTaskService {
	return &ProjectTaskServiceImpl{collection}
}

func (pts *ProjectTaskServiceImpl) GetProjectTask(id string) (*models.ProjectTask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("ProjectTask", err)
	}

	var task *models.ProjectTask

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: taskId}}}}

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "manager"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "manager_info"},
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

	result, err := pts.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer result.Close(ctx)

	if result.Next(ctx) {
		if err := result.Decode(&task); err != nil {
			return nil, &errors.CustomError{
				Message:    "결과 디코딩 오류",
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return task, nil
}

func (pts *ProjectTaskServiceImpl) GetProjectTaskByProject(id string) ([]models.ProjectTask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	projectId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("Project", err)
	}

	var tasks []models.ProjectTask

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "project", Value: projectId}}}}

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "manager"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "manager_info"},
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

	results, err := pts.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &tasks); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return tasks, nil
}

func (pts *ProjectTaskServiceImpl) CreateProjectTask(dto *dto.ProjectTaskCreateDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("dto: %+v", dto)

	task := models.ProjectTask{
		ID:              primitive.NewObjectID(),
		TaskDescription: dto.TaskDescription,
		Status:          dto.Status,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	var err error

	if task.Project, err = utils.ConvertToObjectId(dto.Project); err != nil {
		return utils.ConvertError("Project", err)
	}

	if task.User, err = utils.ConvertToObjectId(dto.Manager); err != nil {
		return utils.ConvertError("User", err)
	}

	if task.Department, err = utils.ConvertToObjectId(dto.Department); err != nil {
		return utils.ConvertError("Department", err)
	}

	fmt.Printf("task: %+v", task)

	_, err = pts.collection.InsertOne(ctx, task)

	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}

func (pts *ProjectTaskServiceImpl) UpdateProjectTask(id string, dto *dto.ProjectTaskUpdateDTO) (*models.ProjectTask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("ProjectTask", err)
	}

	task := bson.M{
		"updated_at": time.Now(),
	}

	if dto.TaskDescription != "" {
		task["task_description"] = dto.TaskDescription
	}

	if dto.Status != "" {
		task["status"] = dto.Status
	}

	if dto.Manager != "" {
		if task["manager"], err = utils.ConvertToObjectId(dto.Manager); err != nil {
			return nil, utils.ConvertError("User", err)
		}
	}

	if dto.Department != "" {
		if task["department"], err = utils.ConvertToObjectId(dto.Department); err != nil {
			return nil, utils.ConvertError("Department", err)
		}
	}

	filter := bson.M{"_id": taskId}
	update := bson.M{"$set": task}

	fmt.Printf("task: %+v", task)

	result := pts.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, &errors.CustomError{
				Message:    "Project Task를 찾을 수 없음",
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

	var updatedTask *models.ProjectTask
	if err := result.Decode(&updatedTask); err != nil {
		return nil, &errors.CustomError{
			Message:    "결과 디코딩 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return updatedTask, nil
}

func (pts *ProjectTaskServiceImpl) DeleteProjectTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return utils.ConvertError("ProjectTask", err)
	}

	filter := bson.M{"_id": taskId}

	result, err := pts.collection.DeleteOne(ctx, filter)
	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if result.DeletedCount == 0 {
		return &errors.CustomError{
			Message:    "Project Task를 찾을 수 없음",
			StatusCode: http.StatusNotFound,
			Err:        mongo.ErrNoDocuments,
		}
	}

	return nil
}
