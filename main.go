package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin 엔진 생성
	r := gin.Default()

	// 루트 엔드포인트에 "Hello, World!" 응답 반환
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 웹 서버 시작
	r.Run(":8080")
}
