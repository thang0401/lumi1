package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// RegisterRoutes đăng ký các route API
func RegisterRoutes(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	r.GET("/data", handleGetData(db, redisClient))
	// r.POST("/data", handlePostData(db, redisClient))
}
