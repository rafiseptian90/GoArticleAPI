package database

import (
	models2 "github.com/rafiseptian90/GoArticle/models"
	"gorm.io/gorm"
	"log"
)

func InitMigration(DB *gorm.DB) {
	err := DB.AutoMigrate(&models2.Tag{}, &models2.Article{}, models2.User{}, models2.Profile{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
