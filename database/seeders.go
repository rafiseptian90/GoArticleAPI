package database

import (
	"github.com/rafiseptian90/GoArticle/database/fakers"
	"gorm.io/gorm"
)

func InitSeeder(DB *gorm.DB) {
	fakers.NewUserSeeders(DB)
	fakers.NewTagSeeders(DB)
	fakers.NewArticleSeeders(DB)
}
