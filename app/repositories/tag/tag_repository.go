package tag

import "github.com/rafiseptian90/GoArticle/app/models"

type TagRepository interface {
	GetTags() []models.Tag
	GetTag(slug string) (models.Tag, error)
	StoreTag(tagRequest *models.TagRequest) error
	UpdateTag(slug string, tagRequest *models.TagRequest) error
	DeleteTag(slug string) error
}
