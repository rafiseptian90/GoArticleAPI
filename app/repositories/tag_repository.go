package repositories

import (
	"errors"
	"github.com/rafiseptian90/GoArticle/app/handlers/requests"
	"github.com/rafiseptian90/GoArticle/app/handlers/responses"
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
)

type TagRepositoryInterface interface {
	GetTags() []responses.TagResponse
	GetTag(tagID int) (responses.TagResponse, error)
	StoreTag(tagRequest *requests.TagRequest) error
	UpdateTag(tagID int, tagRequest *requests.TagRequest) error
	DeleteTag(tagID int) error
}

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(DB *gorm.DB) *TagRepository {
	return &TagRepository{DB: DB}
}

func (repository *TagRepository) GetTags() []responses.TagResponse {
	var categories []responses.TagResponse

	repository.DB.Find(&categories)

	return categories
}

func (repository *TagRepository) GetTag(tagID int) (responses.TagResponse, error) {
	var tag responses.TagResponse

	result := repository.DB.First(&tag, tagID)
	if result.RowsAffected < 1 {
		return tag, errors.New("Tag not found")
	}

	return tag, nil
}

func (repository *TagRepository) StoreTag(tagRequest *requests.TagRequest) error {
	result := repository.DB.Create(&tagRequest)

	if result.RowsAffected < 1 {
		return errors.New("Can't create a new tag")
	}

	return nil
}

func (repository *TagRepository) UpdateTag(tagID int, tagRequest *requests.TagRequest) error {
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
