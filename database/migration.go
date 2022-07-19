package database

import (
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
	"log"
)

func InitMigration(DB *gorm.DB) {
	err := DB.AutoMigrate(&models.Tag{}, &models.Article{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
