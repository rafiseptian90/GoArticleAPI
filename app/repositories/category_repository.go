package repositories

import (
	"errors"
	"github.com/rafiseptian90/GoArticle/app/handlers/requests"
	"github.com/rafiseptian90/GoArticle/app/handlers/responses"
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	GetCategories() []responses.CategoryResponse
	GetCategory(categoryID int) (responses.CategoryResponse, error)
	StoreCategory(categoryRequest *requests.CategoryRequest) error
	UpdateCategory(categoryID int, categoryRequest *requests.CategoryRequest) error
	DeleteCategory(categoryID int) error
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: DB}
}

func (repository *CategoryRepository) GetCategories() []responses.CategoryResponse {
	var categories []responses.CategoryResponse

	repository.DB.Find(&categories)

	return categories
}

func (repository *CategoryRepository) GetCategory(categoryID int) (responses.CategoryResponse, error) {
	var category responses.CategoryResponse

	result := repository.DB.First(&category, categoryID)
	if result.RowsAffected < 1 {
		return category, errors.New("Category not found")
	}

	return category, nil
}

func (repository *CategoryRepository) StoreCategory(categoryRequest *requests.CategoryRequest) error {
	result := repository.DB.Create(&categoryRequest)

	if result.RowsAffected < 1 {
		return errors.New("Can't create a new category")
	}

	return nil
}

func (repository *CategoryRepository) UpdateCategory(categoryID int, categoryRequest *requests.CategoryRequest) error {
	if result := repository.DB.Where("id = ?", categoryID).Updates(categoryRequest); result.RowsAffected < 1 {
		return errors.New("Category is not found")
	}

	return nil
}

func (repository *CategoryRepository) DeleteCategory(categoryID int) error {
	if result := repository.DB.Delete(&models.Category{}, categoryID); result.RowsAffected < 1 {
		return errors.New("Category is not found")
	}

	return nil
}
