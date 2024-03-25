package main

import (
	"context"
	"log"

	"github.com/Kim-DaeHan/all-note-golang/database"
	"github.com/Kim-DaeHan/all-note-golang/routes"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	// .env 파일 로딩
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	//run database
	db := database.ConnectDB()
	defer db.Disconnect(context.Background()) // 애플리케이션이 종료되기 전에 연결 닫기

	routes.SetDependency(db)
	routes.SetupRoutes(router)

	router.Run("localhost:8080")
}
