package main

import (
	"context"
	"log"

	"github.com/Kim-DaeHan/all-note-golang/database"
	"github.com/Kim-DaeHan/all-note-golang/routes"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	// .env 파일 로딩
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// CORS 설정
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	// config.AllowAllOrigins = true // 모든 오리진 허용
	config.AllowCredentials = true
	router.Use(cors.New(config))

	//run database
	db := database.ConnectDB()
	defer db.Disconnect(context.Background()) // 애플리케이션이 종료되기 전에 연결 닫기

	routes.SetDependency(db)
	routes.SetupRoutes(router)

	router.Run("localhost:8080")
}
