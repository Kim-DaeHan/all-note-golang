package main

import (
	"log"

	"github.com/Kim-DaeHan/all-note-golang/cmd"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
