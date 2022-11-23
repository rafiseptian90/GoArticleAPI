package repository

import (
	"encoding/json"
	"errors"
	"github.com/rafiseptian90/GoArticle/models"
	"gorm.io/gorm"
	"log"
)

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(DB *gorm.DB) *TagRepository {
	return &TagRepository{DB: DB}
}

func (repository *TagRepository) GetTags() []models.Tag {
	var tags []models.Tag

	repository.DB.Find(&tags)

	return tags
}

func (repository *TagRepository) GetTag(slug string) (models.Tag, error) {
	var tag models.Tag

	result := repository.DB.Where("slug = ?", slug).First(&tag)
	if result.RowsAffected < 1 {
		return tag, errors.New(result.Error.Error())
	}

	return tag, nil
}

func (repository *TagRepository) StoreTag(tagRequest *models.TagRequest) error {
	var tag map[string]interface{}
	data, _ := json.Marshal(tagRequest)
	if err := json.Unmarshal(data, &tag); err != nil {
		log.Fatalf("Error : %v", err.Error())
	}

	result := repository.DB.Model(&models.Tag{}).Create(tag)

	if result.RowsAffected < 1 {
		return errors.New(result.Error.Error())
	}

	return nil
}

func (repository *TagRepository) UpdateTag(slug string, tagRequest *models.TagRequest) error {
	var tag map[string]interface{}
	data, _ := json.Marshal(tagRequest)
	json.Unmarshal(data, &tag)

	if result := repository.DB.Model(&models.Tag{}).Where("slug = ?", slug).Updates(tag); result.RowsAffected < 1 {
		return errors.New("tag is not found")
	}

	return nil
}

func (repository *TagRepository) DeleteTag(slug string) error {
	if result := repository.DB.Where("slug = ?", slug).Delete(&models.Tag{}); result.RowsAffected < 1 {
		return result.Error
	}

	return nil
}
