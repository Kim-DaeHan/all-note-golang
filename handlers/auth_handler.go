package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/Kim-DaeHan/all-note-golang/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService services.UserService
}

func NewAuthController(userService services.UserService) AuthHandler {
	return AuthHandler{userService}
}

func (ac *AuthHandler) GoogleOAuth(c *gin.Context) {
	code := c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	// Use the code to get the id and access tokens
	tokenRes, err := utils.GetGoogleOauthToken(code)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	fmt.Printf("user: %+v", user)

	resBody := &dto.UserUpdateDTO{
		Email:    user.Email,
		UserName: user.Name,
		Photo:    user.Picture,
		Provider: "google",
		GoogleID: user.Id,
		Verified: &user.Verified_email,
	}

	updatedUser, err := ac.userService.UpsertUser(resBody)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	accessExpiredInStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	accessTokenExpiredIn, err := time.ParseDuration(accessExpiredInStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	refreshExpiredInStr := os.Getenv("REFRESH_TOKEN_EXPIRED_IN")
	refreshTokenExpiredIn, err := time.ParseDuration(refreshExpiredInStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	// Generate Tokens
	access_token, err := utils.CreateToken(accessTokenExpiredIn, updatedUser.ID.Hex(), os.Getenv("ACCESS_TOKEN_JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(refreshTokenExpiredIn, updatedUser.ID.Hex(), os.Getenv("REFRESH_TOKEN_JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	accessTokenMaxAgeStr := os.Getenv("ACCESS_TOKEN_MAXAGE")
	accessTokenMaxAge, err := strconv.Atoi(accessTokenMaxAgeStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	refreshTokenMaxAgeStr := os.Getenv("REFRESH_TOKEN_MAXAGE")
	refreshTokenMaxAge, err := strconv.Atoi(refreshTokenMaxAgeStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	c.SetCookie("access_token", access_token, accessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh_token, refreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", accessTokenMaxAge*60, "/", "localhost", false, false)

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprint(os.Getenv("CLIENT_ORIGIN"), pathUrl))
}

func (ac *AuthHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
