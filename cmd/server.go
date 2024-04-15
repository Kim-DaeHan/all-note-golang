package cmd

import (
	"context"

	"github.com/Kim-DaeHan/all-note-golang/database"
	"github.com/Kim-DaeHan/all-note-golang/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	router := gin.Default()

	// CORS 설정
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	// config.AllowAllOrigins = true // 모든 오리진 허용
	config.AllowCredentials = true
	router.Use(cors.New(config))

	return router
}

func RunServer(router *gin.Engine) {
	//run database
	db := database.ConnectDB()
	defer db.Disconnect(context.Background()) // 애플리케이션이 종료되기 전에 연결 닫기

	routes.SetDependency(db)
	routes.SetupRoutes(router)

	router.Run("localhost:8080")
}
