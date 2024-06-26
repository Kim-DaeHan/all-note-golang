package main

import (
	"log"

	"github.com/Kim-DaeHan/all-note-golang/cmd"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title All Note API
// @version 1.0
// @description all note 어플리케이션 api

// @contact.name API Support
// @contact.url https://github.com/Kim-DaeHan/all-note-golang
// @contact.email kjs50458281@gmail.com

// @BasePath /api
func main() {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	// .env 파일 로딩
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := cmd.SetupServer()
	cmd.RunServer(router)
}
