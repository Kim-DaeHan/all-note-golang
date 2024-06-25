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

type ProjectServiceImpl struct {
	collection *mongo.Collection
}

func NewProjectServiceImpl(collection *mongo.Collection) services.ProjectService {
	return &ProjectServiceImpl{collection}
}

func (ps *ProjectServiceImpl) GetAllProject() ([]models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var projects []models.Project

	results, err := ps.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer results.Close(ctx)

	if err = results.All(ctx, &projects); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return projects, nil
}

func (ps *ProjectServiceImpl) CreateProject(dto *dto.ProjectCreateDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("dto: %+v", dto)

	project := models.Project{
		ID:        primitive.NewObjectID(),
		Name:      dto.Name,
		StartDt:   dto.StartDt,
		EndDt:     dto.EndDt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Printf("project: %+v", project)

	_, err := ps.collection.InsertOne(ctx, project)

	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}

func (ps *ProjectServiceImpl) UpdateProject(id string, dto *dto.ProjectUpdateDTO) (*models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	projectId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("Project", err)
	}

	project := bson.M{
		"updated_at": time.Now(),
	}

	if dto.Name != "" {
		project["name"] = dto.Name
	}

	if !dto.StartDt.IsZero() {
		project["start_dt"] = dto.StartDt
	}

	if !dto.EndDt.IsZero() {
		project["end_dt"] = dto.EndDt
	}

	filter := bson.M{"_id": projectId}
	update := bson.M{"$set": project}

	fmt.Printf("project: %+v", project)

	result := ps.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, &errors.CustomError{
				Message:    "Project를 찾을 수 없음",
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

	var updatedProject *models.Project
	if err := result.Decode(&updatedProject); err != nil {
		return nil, &errors.CustomError{
			Message:    "결과 디코딩 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return updatedProject, nil
}

func (ps *ProjectServiceImpl) DeleteProject(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	projectId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return utils.ConvertError("Project", err)
	}

	filter := bson.M{"_id": projectId}

	result, err := ps.collection.DeleteOne(ctx, filter)
	if err != nil {
		return &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if result.DeletedCount == 0 {
		return &errors.CustomError{
			Message:    "Project를 찾을 수 없음",
			StatusCode: http.StatusNotFound,
			Err:        mongo.ErrNoDocuments,
		}
	}

	return nil
}
