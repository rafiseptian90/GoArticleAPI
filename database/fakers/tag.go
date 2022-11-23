package fakers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gosimple/slug"
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
)

func NewTagSeeders(DB *gorm.DB) {
	for i := 1; i <= 7; i++ {
		word := faker.Word()

		DB.Create(&models.Tag{Slug: slug.Make(word), Name: word})
	}
}
