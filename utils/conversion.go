package utils

import (
	"fmt"
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConvertToObjectId converts a hex string to a primitive.ObjectID.
func ConvertToObjectId(hex string) (primitive.ObjectID, error) {
	if hex == "" {
		return primitive.NilObjectID, nil
	}
	return primitive.ObjectIDFromHex(hex)
}

// CustomError creates a new CustomError for ObjectID conversion errors.
func ConvertError(fieldName string, err error) *errors.CustomError {
	return &errors.CustomError{
		Message:    fmt.Sprintf("%s ObjectID 변환 오류", fieldName),
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
}
