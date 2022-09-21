package fakers

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func NewArticleTagSeeder(DB *gorm.DB) {
	var article_tags []map[string]interface{}

	for i := 1; i <= 10; i++ {
		articleID, _ := faker.RandomInt(1, 50, 1)
		tagID, _ := faker.RandomInt(1, 7, 1)

		article_tags = append(article_tags, map[string]interface{}{
			"article_id": articleID,
			"tag_id":     tagID,
		})
	}

	DB.Table("article_tags").Create(article_tags)
}
