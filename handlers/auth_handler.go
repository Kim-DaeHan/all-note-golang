package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/Kim-DaeHan/all-note-golang/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService services.UserService
}

func NewAuthHandler(userService services.UserService) AuthHandler {
	return AuthHandler{userService}
}

func (ah *AuthHandler) GoogleOAuth(ctx *gin.Context) {
	code := ctx.Query("code")
	var pathUrl string = "/"

	fullURL := ctx.Request.URL.String()

	fmt.Println("code: ", code)
	fmt.Println("fullUrl: ", fullURL)

	if ctx.Query("state") != "" {
		pathUrl = ctx.Query("state")
	}

	if code == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	// Use the code to get the id and access tokens
	tokenRes, err := utils.GetGoogleOauthToken(code)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
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

	updatedUser, err := ah.userService.UpsertUser(resBody)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	accessExpiredInStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	accessTokenExpiredIn, err := time.ParseDuration(accessExpiredInStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	refreshExpiredInStr := os.Getenv("REFRESH_TOKEN_EXPIRED_IN")
	refreshTokenExpiredIn, err := time.ParseDuration(refreshExpiredInStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	// Generate Tokens
	access_token, err := utils.CreateToken(accessTokenExpiredIn, updatedUser.ID.Hex(), os.Getenv("ACCESS_TOKEN_JWT_SECRET"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(refreshTokenExpiredIn, updatedUser.ID.Hex(), os.Getenv("REFRESH_TOKEN_JWT_SECRET"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	accessTokenMaxAgeStr := os.Getenv("ACCESS_TOKEN_MAXAGE")
	accessTokenMaxAge, err := strconv.Atoi(accessTokenMaxAgeStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	refreshTokenMaxAgeStr := os.Getenv("REFRESH_TOKEN_MAXAGE")
	refreshTokenMaxAge, err := strconv.Atoi(refreshTokenMaxAgeStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	ctx.SetCookie("access_token", access_token, accessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, refreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", accessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprint(os.Getenv("CLIENT_ORIGIN"), pathUrl))
}

func (ah *AuthHandler) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ah *AuthHandler) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	fmt.Printf("user: %+v", currentUser)
	ctx.JSON(http.StatusOK, currentUser)

}

func (ah *AuthHandler) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	sub, err := utils.ValidateToken(cookie, os.Getenv("REFRESH_TOKEN_JWT_SECRET"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := ah.userService.GetUser(sub.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	accessExpiredInStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	accessTokenExpiredIn, err := time.ParseDuration(accessExpiredInStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	accessTokenMaxAgeStr := os.Getenv("ACCESS_TOKEN_MAXAGE")
	accessTokenMaxAge, err := strconv.Atoi(accessTokenMaxAgeStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	access_token, err := utils.CreateToken(accessTokenExpiredIn, user.ID.Hex(), os.Getenv("ACCESS_TOKEN_JWT_SECRET"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, accessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", accessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}
