package repositories

import (
	"encoding/json"
	"errors"
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
)

type ArticleRepositoryInterface interface {
	GetArticles() map[string]interface{}
	GetArticlesByTags(tags []string) []models.Article
	GetArticle(articleID int) (models.Article, error)
	StoreArticle(articleRequest *models.ArticleRequest) error
	UpdateArticle(articleID int, articleRequest *models.ArticleRequest) error
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

func (repository *ArticleRepository) GetArticles() map[string]interface{} {
	var articles []models.Article

	repository.DB.Order("seen asc").Preload("User.Profile").Preload("Tags").Find(&articles)

	return map[string]interface{}{
		"trending_articles": articles[0:6],
		"articles":          articles[6:],
	}
}

func (repository ArticleRepository) GetArticlesByTags(tags []string) []models.Article {
	var articles []models.Article

	repository.DB.Preload("Tags").Where("EXISTS (SELECT * FROM article_tags WHERE article_tags.article_id = articles.id AND article_tags.tag_id IN (?))", tags).Find(&articles)

	return articles
}

func (repository ArticleRepository) GetArticle(articleID int) (models.Article, error) {
	var article models.Article

	if result := repository.DB.Model(&article).Preload("Tags").First(&article, articleID); result.RowsAffected < 1 {
		return article, errors.New("Article not found")
	}

	return article, nil
}

func (repository ArticleRepository) StoreArticle(articleRequest *models.ArticleRequest) error {
	var article map[string]interface{}
	data, _ := json.Marshal(articleRequest)
	json.Unmarshal(data, &article)

	if result := repository.DB.Model(&models.Article{}).Create(article); result.RowsAffected < 1 {
		return errors.New("Can't create the article")
	}

	return nil
}

func (repository ArticleRepository) UpdateArticle(articleID int, articleRequest *models.ArticleRequest) error {
	var article map[string]interface{}
	data, _ := json.Marshal(articleRequest)
	json.Unmarshal(data, &article)

	if result := repository.DB.Model(&models.Article{}).Where("id = ?", articleID).Updates(article); result.RowsAffected < 1 {
		return errors.New("Article is not found")
	}

	return nil
}

func (repository ArticleRepository) DeleteArticle(articleID int) error {
	if result := repository.DB.Delete(&models.Article{}, articleID); result.RowsAffected < 1 {
		return errors.New("Article is not found")
	}

	return nil
}
