package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func Connect() *gorm.DB {
	dsn := "root:pass123@tcp(localhost:3306)/post-hub-app?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
