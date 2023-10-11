package database

import (
	"PostHubApp/domain/use_case/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {
	dsn := "root:pass123@tcp(localhost:3306)/post-hub-app?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate will create the tables if they don't exist and update their schema if necessary
	err = db.AutoMigrate(&entity.Comment{}, &entity.Post{}, &entity.PostModeration{}, &entity.User{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
