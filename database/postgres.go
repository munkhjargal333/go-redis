package database

import (
	"fmt"
	"log"
	"route-redis/shared"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var RDB *redis.Client

func Connect() {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ulaanbaatar`,
		shared.Config.DB.Host,
		shared.Config.DB.User,
		shared.Config.DB.Password,
		shared.Config.DB.Name,
		shared.Config.DB.Port,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	DB = db
	DB.Logger = logger.Default.LogMode(logger.Info)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	RDB = rdb
}
