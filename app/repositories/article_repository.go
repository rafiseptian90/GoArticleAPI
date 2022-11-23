package repositories

import (
	"encoding/json"
	"errors"
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
	"time"
)

type ArticleRepositoryInterface interface {
	GetArticles() map[string]interface{}
	GetArticlesByTags(tags []string) []models.Article
	GetTrendingArticlesByTags(tags []string) []models.Article
	GetLatestArticlesByTags(tags []string) []models.Article
	GetBestArticlesByTags(tags []string) []models.Article
	GetArticle(articleSlug string) (map[string]interface{}, error)
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

func (repository ArticleRepository) GetArticles() map[string]interface{} {
	var articles []models.Article

	repository.DB.Order("seen desc").Preload("User.Profile").Preload("Tags").Find(&articles)

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

func (repository ArticleRepository) GetTrendingArticlesByTags(tags []string) []models.Article {
	var articles []models.Article

	now := time.Now()
	currentTime := now.Format("2006-01-02 15:04:05")
	lastWeek := now.AddDate(0, 0, -7).Format("2006-01-02 15:04:05")

	repository.DB.Preload("User.Profile").Preload("Tags").Where("published_at BETWEEN ? AND ?", lastWeek, currentTime).Order("seen desc").Find(&articles)

	return articles
}

func (repository ArticleRepository) GetLatestArticlesByTags(tags []string) []models.Article {
	var articles []models.Article

	repository.DB.Preload("User.Profile").Preload("Tags").Where("EXISTS (SELECT * FROM article_tags WHERE article_tags.article_id = articles.id AND article_tags.tag_id IN (?))", tags).Order("published_at desc").Find(&articles)

	return articles
}

func (repository ArticleRepository) GetBestArticlesByTags(tags []string) []models.Article {
	var articles []models.Article

	repository.DB.Preload("User.Profile").Preload("Tags").Where("EXISTS (SELECT * FROM article_tags WHERE article_tags.article_id = articles.id AND article_tags.tag_id IN (?))", tags).Order("seen desc").Find(&articles)

	return articles
}

func (repository ArticleRepository) GetArticle(articleSlug string) (map[string]interface{}, error) {
	var article models.Article
	var popularArticles, recommendedArticles []models.Article

	// Query to get single article by slug
	if result := repository.DB.Model(&article).Preload("User.Profile").Preload("Tags").Where("slug = ?", articleSlug).Find(&article); result.RowsAffected < 1 {
		return nil, result.Error
	}

	// Query to get 5 popular articles based on that article creator
	if result := repository.DB.Where("user_id = ?", article.UserId).Not("id = ?", article.Id).Limit(5).Order("seen desc").Preload("User.Profile").Preload("Tags").Find(&popularArticles); result.RowsAffected < 1 {
		return nil, result.Error
	}

	// Query to get 8 random articles except the articles of that creator
	if result := repository.DB.Not("user_id = ? ", article.UserId).Order("random()").Limit(8).Preload("User.Profile").Preload("Tags").Find(&recommendedArticles); result.RowsAffected < 1 {
		return nil, result.Error
	}

	return map[string]interface{}{
		"article":              article,
		"user_articles":        popularArticles,
		"recommended_articles": recommendedArticles,
	}, nil
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
