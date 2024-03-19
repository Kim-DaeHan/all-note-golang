package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/errors"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

// var validate = validator.New()

func GetAllUser(c *gin.Context) {
	users, err := services.GetAllUser()

	if err != nil {
		// CustomError 인터페이스로 형변환이 성공하면 customErr에는 *errors.CustomError 타입의 값이 할당되고, ok 변수에는 true가 할당
		customErr, ok := err.(*errors.CustomError)
		if ok {
			statusCode := customErr.Status()
			c.JSON(statusCode, gin.H{"message": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "Data": users})
}

// // 다른 컨트롤러 함수들을 여기에 추가할 수 있습니다.
// func CreateUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		var user models.User
// 		defer cancel()

// 		//validate the request body
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		//use the validator library to validate required fields
// 		if validationErr := validate.Struct(&user); validationErr != nil {
// 			c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
// 			return
// 		}

// 		newUser := models.User{
// 			Name:     user.Name,
// 			Location: user.Location,
// 			Title:    user.Title,
// 		}

// 		result, err := config.UserCollection.InsertOne(ctx, newUser)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
// 	}
// }
