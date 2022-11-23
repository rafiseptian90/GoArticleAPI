package database

import (
	fakers2 "github.com/rafiseptian90/GoArticle/internal/database/fakers"
	"gorm.io/gorm"
)

func InitSeeder(DB *gorm.DB) {
	fakers2.NewUserSeeders(DB)
	fakers2.NewTagSeeders(DB)
	fakers2.NewArticleSeeders(DB)
	fakers2.NewArticleTagSeeder(DB)
}
