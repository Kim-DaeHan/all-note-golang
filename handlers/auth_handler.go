package handlers

import (
	"net/http"

	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService services.UserService
}

func NewAuthController(userService services.UserService) AuthHandler {
	return AuthHandler{userService}
}

func (ac *AuthHandler) GoogleOAuth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "successfully", "data": "aaa"})
}
