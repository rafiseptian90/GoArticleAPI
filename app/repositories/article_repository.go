package repositories

import (
	"errors"
	"github.com/rafiseptian90/GoArticle/app/handlers/requests"
	"github.com/rafiseptian90/GoArticle/app/handlers/responses"
	"github.com/rafiseptian90/GoArticle/app/models"
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
	var articles []responses.ArticleResponse

	repository.DB.Find(&articles)

	return articles
}

func (repository *ArticleRepository) GetArticle(articleID int) (responses.ArticleResponse, error) {
	var article responses.ArticleResponse

	if result := repository.DB.First(&article, articleID); result.RowsAffected < 1 {
		return article, errors.New("Article not found")
	}

	return article, nil
}

func (repository *ArticleRepository) StoreArticle(articleRequest *requests.ArticleRequest) error {
	if result := repository.DB.Create(articleRequest); result.RowsAffected < 1 {
		return errors.New("Can't create the article")
	}

	return nil
}

func (repository *ArticleRepository) UpdateArticle(articleID int, articleRequest *requests.ArticleRequest) error {
	if result := repository.DB.Where("id = ?", articleID).Updates(articleRequest); result.RowsAffected < 1 {
		return errors.New("Article is not found")
	}

	return nil
}

func (repository *ArticleRepository) DeleteArticle(articleID int) error {
	if result := repository.DB.Delete(&models.Article{}, articleID); result.RowsAffected < 1 {
		return errors.New("Article is not found")
	}

	return nil
}
