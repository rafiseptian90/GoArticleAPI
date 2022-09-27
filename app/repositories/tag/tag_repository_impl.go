package tag

import (
	"encoding/json"
	"errors"
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewTagRepository(DB *gorm.DB) *Repository {
	return &Repository{DB: DB}
}

func (repository *Repository) GetTags() []models.Tag {
	var tags []models.Tag

	repository.DB.Find(&tags)

	return tags
}

func (repository *Repository) GetTag(slug string) (models.Tag, error) {
	var tag models.Tag

	result := repository.DB.Where("slug = ?", slug).Preload("Articles.User.Profile").First(&tag)
	if result.RowsAffected < 1 {
		return tag, errors.New("Tag not found")
	}

	return tag, nil
}

func (repository *Repository) StoreTag(tagRequest *models.TagRequest) error {
	var tag map[string]interface{}
	data, _ := json.Marshal(tagRequest)
	json.Unmarshal(data, &tag)

	result := repository.DB.Model(&models.Tag{}).Create(tag)

	if result.RowsAffected < 1 {
		return errors.New("Can't create a new tag")
	}

	return nil
}

func (repository *Repository) UpdateTag(slug string, tagRequest *models.TagRequest) error {
	var tag map[string]interface{}
	data, _ := json.Marshal(tagRequest)
	json.Unmarshal(data, &tag)

	if result := repository.DB.Model(&models.Tag{}).Where("slug = ?", slug).Updates(tag); result.RowsAffected < 1 {
		return errors.New("Tag is not found")
	}

	return nil
}

func (repository *Repository) DeleteTag(slug string) error {
	if result := repository.DB.Where("slug = ?", slug).Delete(&models.Tag{}); result.RowsAffected < 1 {
		return result.Error
	}

	return nil
}
