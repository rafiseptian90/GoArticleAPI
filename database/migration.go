package database

import (
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
	"log"
)

func InitMigration(DB *gorm.DB) {
	err := DB.AutoMigrate(&models.Tag{}, &models.Article{}, models.User{}, models.Profile{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
