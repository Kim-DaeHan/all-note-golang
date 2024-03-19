package services

import (
	"context"
	"net/http"
	"time"

	"github.com/Kim-DaeHan/all-note-golang/database"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUser() ([]models.User, error) {
	var DB = database.ConnectDB()
	var userCollection = database.GetCollection(DB, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	results, err := userCollection.Find(ctx, bson.M{})

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
