package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func GetAllUser(c *gin.Context) {
	users, err := services.GetAllUser()

	if err != nil {
		// CustomError 인터페이스로 형변환이 성공하면 customErr에는 *errors.CustomError 타입의 값이 할당되고, ok 변수에는 true가 할당
		customErr, ok := err.(*errors.CustomError)
		if ok {
			statusCode := customErr.Status()
			c.JSON(statusCode, gin.H{"err": customErr.Err.Error(), "message": customErr.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": users})
}

func CreateUser(c *gin.Context) {
	var dto dto.UserCreateDTO

	//validate the request body
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&dto); validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": validationErr.Error()})
		return
	}

	result, err := services.CreateUser(dto)

	if err != nil {
		// CustomError 인터페이스로 형변환이 성공하면 customErr에는 *errors.CustomError 타입의 값이 할당되고, ok 변수에는 true가 할당
		customErr, ok := err.(*errors.CustomError)
		if ok {
			statusCode := customErr.Status()
			c.JSON(statusCode, gin.H{"err": customErr.Err.Error(), "message": customErr.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": result})
}
