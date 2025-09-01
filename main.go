package main

import (
	"fmt"
	"log"
	"lumi/config"
	"lumi/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// Kết nối MariaDB & Redis
	db, redisClient, err := config.InitDB()
	if err != nil {
		log.Fatalf("❌ Lỗi kết nối DB hoặc Redis: %v", err)
	}

	// Khởi tạo router Gin
	r := gin.Default()

	// Đăng ký routes
	controller.RegisterRoutes(r, db, redisClient)

	// Chạy server
	fmt.Println("🚀 Server chạy tại: http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Không thể chạy server: %v", err)
	}
}
