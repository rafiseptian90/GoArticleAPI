package fakers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gosimple/slug"
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
	"math/rand"
	"reflect"
	"time"
)

func NewArticleSeeders(DB *gorm.DB) {
	for i := 1; i <= 15; i++ {
		sentences, _ := faker.GetLorem().Sentence(reflect.Value{})
		paragraph, _ := faker.GetLorem().Paragraph(reflect.Value{})

		DB.Create(&models.Article{
			UserId:      uint(rand.Intn(5-1) + 1),
			Slug:        slug.Make(sentences.(string)),
			Title:       sentences.(string),
			Content:     paragraph.(string),
			Seen:        uint(rand.Intn(1000-1) + 1),
			PublishedAt: time.Now(),
		})
	}
}
