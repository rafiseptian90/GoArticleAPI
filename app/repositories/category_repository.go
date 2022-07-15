package repositories

import (
	"github.com/rafiseptian90/GoArticle/app/handlers/requests"
	"github.com/rafiseptian90/GoArticle/app/handlers/responses"
)

type CategoryRepositoryInterface interface {
	GetCategories() []responses.CategoryResponse
	GetCategory(categoryID int) (responses.CategoryResponse, error)
	StoreCategory(categoryRequest *requests.CategoryRequest) error
	UpdateCategory(categoryID int, categoryRequest *requests.CategoryRequest) error
	DeleteCategory(categoryID int) error
}

type CategoryRepository struct {
}

func (repository *CategoryRepository) GetCategories() []responses.CategoryResponse {
	//TODO implement me
	panic("implement me")
}

func (repository *CategoryRepository) GetCategory(categoryID int) (responses.CategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *CategoryRepository) StoreCategory(categoryRequest *requests.CategoryRequest) error {
	//TODO implement me
	panic("implement me")
}

func (repository *CategoryRepository) UpdateCategory(categoryID int, categoryRequest *requests.CategoryRequest) error {
	//TODO implement me
	panic("implement me")
}

func (repository *CategoryRepository) DeleteCategory(categoryID int) error {
	//TODO implement me
	panic("implement me")
}
