package repositories

import (
	"github.com/rafiseptian90/GoArticle/app/handlers/requests"
	"github.com/rafiseptian90/GoArticle/app/handlers/responses"
	"gorm.io/gorm"
)

type ArticleRepositoryInterface interface {
	GetArticles() []responses.ArticleResponse
	GetArticle(articleID int) (responses.ArticleResponse, error)
	StoreArticle(articleRequest *requests.ArticleRequest) error
	UpdateArticle(articleID int, articleRequest *requests.ArticleRequest) error
	DeleteArticle(articleID int) error
}

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(DB *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		DB: DB,
	}
}

func (repository *ArticleRepository) GetArticles() []responses.ArticleResponse {
	//TODO implement me
	panic("implement me")
}

func (repository *ArticleRepository) GetArticle(articleID int) (responses.ArticleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *ArticleRepository) StoreArticle(articleRequest *requests.ArticleRequest) error {
	//TODO implement me
	panic("implement me")
}

func (repository *ArticleRepository) UpdateArticle(articleID int, articleRequest *requests.ArticleRequest) error {
	//TODO implement me
	panic("implement me")
}

func (repository *ArticleRepository) DeleteArticle(articleID int) error {
	//TODO implement me
	panic("implement me")
}
