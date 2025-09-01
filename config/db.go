package config

import (
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

// InitDB khởi tạo kết nối MariaDB và Redis
func InitDB() (*gorm.DB, *redis.Client, error) {
	// Thông tin kết nối DB
	username := "lumi"
	password := "nzb83PcrWtkMyfPf"
	host := "103.90.225.131"
	port := "3306"
	dbName := "lumi"

	// DSN MariaDB
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Kết nối MariaDB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	// Connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Kết nối Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	DB = db
	RedisClient = redisClient
	return db, redisClient, nil
}
