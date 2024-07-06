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

type JobApplicationServiceImpl struct {
	collection *mongo.Collection
}

func NewJobApplicationServiceImpl(collection *mongo.Collection) services.JobApplicationService {
	return &JobApplicationServiceImpl{collection}
}

func (js *JobApplicationServiceImpl) GetAllJobApplication() ([]models.JobApplication, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var jobApplications []models.JobApplication

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "manager"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "manager_info"},
	}}}

	lookupDepartmentStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "departments"},
		{Key: "localField", Value: "department"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "department_info"},
	}}}

	pipeline := mongo.Pipeline{lookupUserStage, lookupDepartmentStage}

	results, err := js.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &jobApplications); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return jobApplications, nil
}

func (js *JobApplicationServiceImpl) GetJobApplication(id string) (*models.JobApplication, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobApplicationId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("JobApplication", err)
	}

	var jobApplication *models.JobApplication

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: jobApplicationId}}}}

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "manager"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "manager_info"},
	}}}

	lookupDepartmentStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "departments"},
		{Key: "localField", Value: "department"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "department_info"},
	}}}

	pipeline := mongo.Pipeline{matchStage, lookupUserStage, lookupDepartmentStage}

	result, err := js.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer result.Close(ctx)

	if result.Next(ctx) {
		if err := result.Decode(&jobApplication); err != nil {
			return nil, &errors.CustomError{
				Message:    "결과 디코딩 오류",
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return jobApplication, nil
}

func (js *JobApplicationServiceImpl) GetJobApplicationByManager(id string) ([]models.JobApplication, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	managerId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("User", err)
	}

	var jobApplications []models.JobApplication

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "manager", Value: managerId}}}}

	lookupUserStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "users"},
		{Key: "localField", Value: "manager"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "manager_info"},
	}}}

	lookupDepartmentStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "departments"},
		{Key: "localField", Value: "department"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "department_info"},
	}}}

	pipeline := mongo.Pipeline{matchStage, lookupUserStage, lookupDepartmentStage}

	results, err := js.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &jobApplications); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return jobApplications, nil
}

func (js *JobApplicationServiceImpl) CreateJobApplication(dto *dto.JobApplicationCreateDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("dto: %+v", dto)

	jobApplication := models.JobApplication{
		ID:            primitive.NewObjectID(),
		ApplicantName: dto.ApplicantName,
		Position:      dto.Position,
		Task:          dto.Task,
		Stage:         dto.Stage,
		Location:      dto.Location,
		StartDt:       dto.StartDt,
		EndDt:         dto.EndDt,
		Status:        dto.Status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	var err error

	if jobApplication.User, err = utils.ConvertToObjectId(dto.User); err != nil {
		return utils.ConvertError("User", err)
	}

	if jobApplication.Department, err = utils.ConvertToObjectId(dto.Department); err != nil {
		return utils.ConvertError("Department", err)
	}

	fmt.Printf("jobApplication: %+v", jobApplication)

	_, err = js.collection.InsertOne(ctx, jobApplication)

	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}

func (js *JobApplicationServiceImpl) UpdateJobApplication(id string, dto *dto.JobApplicationUpdateDTO) (*models.JobApplication, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobApplicationId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("JobApplication", err)
	}

	jobApplication := bson.M{
		"updated_at": time.Now(),
	}

	if dto.ApplicantName != "" {
		jobApplication["applicant_name"] = dto.ApplicantName
	}

	if dto.Position != "" {
		jobApplication["position"] = dto.Position
	}

	if dto.Task != "" {
		jobApplication["task"] = dto.Task
	}

	if dto.Stage != "" {
		jobApplication["stage"] = dto.Stage
	}

	if dto.Location != "" {
		jobApplication["location"] = dto.Location
	}

	if dto.Status != "" {
		jobApplication["status"] = dto.Status
	}

	if !dto.StartDt.IsZero() {
		jobApplication["start_dt"] = dto.StartDt
	}

	if !dto.EndDt.IsZero() {
		jobApplication["end_dt"] = dto.EndDt
	}

	if dto.Department != "" {
		if jobApplication["department"], err = utils.ConvertToObjectId(dto.Department); err != nil {
			return nil, utils.ConvertError("Department", err)
		}
	}

	if dto.User != "" {
		if jobApplication["manager"], err = utils.ConvertToObjectId(dto.User); err != nil {
			return nil, utils.ConvertError("User", err)
		}
	}

	filter := bson.M{"_id": jobApplicationId}
	update := bson.M{"$set": jobApplication}

	fmt.Printf("jobApplication: %+v", jobApplication)

	result := js.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, &errors.CustomError{
				Message:    "JobApplication을 찾을 수 없음",
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

	var updatedJobApplication *models.JobApplication
	if err := result.Decode(&updatedJobApplication); err != nil {
		return nil, &errors.CustomError{
			Message:    "결과 디코딩 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return updatedJobApplication, nil
}

func (js *JobApplicationServiceImpl) DeleteJobApplication(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jobApplicationId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return utils.ConvertError("JobApplication", err)
	}

	filter := bson.M{"_id": jobApplicationId}

	result, err := js.collection.DeleteOne(ctx, filter)
	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if result.DeletedCount == 0 {
		return &errors.CustomError{
			Message:    "JobApplication을 찾을 수 없음",
			StatusCode: http.StatusNotFound,
			Err:        mongo.ErrNoDocuments,
		}
	}

	return nil
}
