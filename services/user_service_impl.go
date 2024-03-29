package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserServiceImpl struct {
	collection *mongo.Collection
}

func NewUserServiceImpl(collection *mongo.Collection) UserService {
	return &UserServiceImpl{collection}
}

func (us *UserServiceImpl) GetAllUser() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User

	query := bson.M{}
	results, err := us.collection.Find(ctx, query)
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

	if err = results.All(ctx, &users); err != nil {
		return nil, &errors.CustomError{
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

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &errors.CustomError{
			Message:    "잘못된 ID 형식",
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	var users *models.User

	query := bson.M{"_id": objID}

	err = us.collection.FindOne(ctx, query).Decode(&users)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return users, nil
}

func (us *UserServiceImpl) CreateUser(dto dto.UserCreateDTO) (*mongo.InsertOneResult, error) {
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

	result, err := us.collection.InsertOne(ctx, user)

	if err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return result, nil
}

func (us *UserServiceImpl) UpsertUser(dto *dto.UserUpdateDTO) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := bson.M{
		"google_id":  dto.GoogleID,
		"email":      dto.Email,
		"user_name":  dto.UserName,
		"verified":   dto.Verified,
		"provider":   dto.Provider,
		"photo":      dto.Photo,
		"updated_at": time.Now(),
	}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(1)
	query := bson.D{{Key: "email", Value: dto.Email}}
	update := bson.D{{Key: "$set", Value: user}, {Key: "$setOnInsert", Value: bson.M{"create_at": time.Now()}}}
	res := us.collection.FindOneAndUpdate(ctx, query, update, opts)

	var updatedUser *models.User

	if err := res.Decode(&updatedUser); err != nil {
		return nil, &errors.CustomError{
			Message:    "내부 서버 오류",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return updatedUser, nil
}
