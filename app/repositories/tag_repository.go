package repositories

import (
	"errors"
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
)

type TagRepositoryInterface interface {
	GetTags() []models.Tag
	GetTag(tagID int) (models.Tag, error)
	StoreTag(tagRequest *models.Tag) error
	UpdateTag(tagID int, tagRequest *models.Tag) error
	DeleteTag(tagID int) error
}

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

func (repository *TagRepository) GetTag(tagID int) (models.Tag, error) {
	var tag models.Tag

	result := repository.DB.First(&tag, tagID)
	if result.RowsAffected < 1 {
		return tag, errors.New("Tag not found")
	}

	return tag, nil
}

func (repository *TagRepository) StoreTag(tagRequest *models.Tag) error {
	result := repository.DB.Create(&tagRequest)

	if result.RowsAffected < 1 {
		return errors.New("Can't create a new tag")
	}

	return nil
}

func (repository *TagRepository) UpdateTag(tagID int, tagRequest *models.Tag) error {
	if result := repository.DB.Where("id = ?", tagID).Updates(tagRequest); result.RowsAffected < 1 {
		return errors.New("Tag is not found")
	}

	return nil
}

func (repository *TagRepository) DeleteTag(tagID int) error {
	if result := repository.DB.Delete(&models.Tag{}, tagID); result.RowsAffected < 1 {
		return errors.New("Tag is not found")
	}

	return nil
}
