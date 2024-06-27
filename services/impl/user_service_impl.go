package impl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	anErr "github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/models"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/Kim-DaeHan/all-note-golang/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserServiceImpl struct {
	collection *mongo.Collection
}

func NewUserServiceImpl(collection *mongo.Collection) services.UserService {
	return &UserServiceImpl{collection}
}

func (us *UserServiceImpl) GetAllUser() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User

	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "departments"},
		{Key: "localField", Value: "department"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "department_info"},
	}}}

	pipeline := mongo.Pipeline{lookupStage}

	// query := bson.M{}
	// results, err := us.collection.Find(ctx, query)
	results, err := us.collection.Aggregate(ctx, pipeline)
	// err := us.collection.FindOne(ctx, query).Decode(&users)

	if err != nil {
		return nil, &anErr.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)

	if err = results.All(ctx, &users); err != nil {
		return nil, &anErr.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return users, nil
}

func (us *UserServiceImpl) GetUser(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, utils.ConvertError("User", err)
	}

	var users *models.User

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: userId}}}}

	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "departments"},
		{Key: "localField", Value: "department"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "department_info"},
	}}}

	pipeline := mongo.Pipeline{matchStage, lookupStage}

	result, err := us.collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, &anErr.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer result.Close(ctx)

	if result.Next(ctx) {
		fmt.Println("result: ", result)
		if err := result.Decode(&users); err != nil {
			return nil, &anErr.CustomError{
				Message:    "결과 디코딩 오류",
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return users, nil
}

func (us *UserServiceImpl) CreateUser(dto *dto.UserCreateDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("dto: %+v", dto)

	user := models.User{
		ID:        primitive.NewObjectID(),
		GoogleID:  dto.GoogleID,
		Email:     dto.Email,
		UserName:  dto.UserName,
		Verified:  dto.Verified,
		Provider:  dto.Provider,
		Photo:     dto.Photo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Printf("user: %+v", user)

	_, err := us.collection.InsertOne(ctx, user)

	if err != nil {
		return &anErr.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}

func (us *UserServiceImpl) UpsertUser(dto *dto.UserUpdateDTO) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error

	if dto.Email == "" {
		return nil, errors.New("email cannot be empty")
	}

	user := bson.M{
		"email":      dto.Email,
		"updated_at": time.Now(),
	}

	if dto.GoogleID != "" {
		user["google_id"] = dto.GoogleID
	}

	if dto.UserName != "" {
		user["user_name"] = dto.UserName
	}

	if dto.Verified != nil {
		user["verified"] = dto.Verified
	}

	if dto.Provider != "" {
		user["provider"] = dto.Provider
	}

	if dto.Photo != "" {
		user["photo"] = dto.Photo
	}

	if dto.Position != "" {
		user["position"] = dto.Position
	}

	if dto.Department != "" {
		if user["department"], err = utils.ConvertToObjectId(dto.Department); err != nil {
			return nil, utils.ConvertError("Department", err)
		}
	}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(1)
	query := bson.D{{Key: "email", Value: dto.Email}}
	update := bson.D{{Key: "$set", Value: user}, {Key: "$setOnInsert", Value: bson.M{"create_at": time.Now()}}}
	result := us.collection.FindOneAndUpdate(ctx, query, update, opts)

	var updatedUser *models.User

	// if{} 안에서만 사용가능한 err 변수
	if err := result.Decode(&updatedUser); err != nil {
		return nil, &anErr.CustomError{
			Message:    "결과 디코딩 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return updatedUser, nil
}
