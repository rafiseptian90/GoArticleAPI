package testing

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafiseptian90/GoArticle/app/controllers"
	"github.com/rafiseptian90/GoArticle/app/handlers/requests"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	"github.com/rafiseptian90/GoArticle/config"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewArticleTest() (*gin.Engine, *controllers.ArticleController) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := config.DBConnection()
	router := gin.Default()
	articleRepository := repositories.NewArticleRepository(DB)
	articleController := controllers.NewArticleController(articleRepository)

	return router, articleController
}

func TestGetArticles(t *testing.T) {
	router, articleController := NewArticleTest()

	router.GET("/article", articleController.Index)

	request, _ := http.NewRequest(http.MethodGet, "/article", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetArticle(t *testing.T) {
	router, articleController := NewArticleTest()

	router.GET("/article/:articleID", articleController.Index)

	t.Run("It should take one article by articleID", func(t *testing.T) {
		articleID := "1"

		request, _ := http.NewRequest(http.MethodGet, "/article/"+articleID, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return article not found", func(t *testing.T) {
		articleID := "999"

		request, _ := http.NewRequest(http.MethodGet, "/article/"+articleID, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func TestStoreArticle(t *testing.T) {
	router, articleController := NewArticleTest()

	router.POST("/article", articleController.Store)

	t.Run("It should create a new article", func(t *testing.T) {
		articleRequest := requests.ArticleRequest{
			Title:   "A new article",
			Content: "New article content",
		}
		requestBody, _ := json.Marshal(articleRequest)

		request, _ := http.NewRequest(http.MethodPost, "/article", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return a bad request error", func(t *testing.T) {
		articleRequest := requests.ArticleRequest{}
		requestBody, _ := json.Marshal(articleRequest)

		request, _ := http.NewRequest(http.MethodPost, "/article", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestUpdateArticle(t *testing.T) {
	router, articleController := NewArticleTest()

	router.PUT("/article/:articleID", articleController.Update)

	t.Run("It should update an article by articleID", func(t *testing.T) {
		articleID := "1"
		articleRequest := requests.ArticleRequest{
			Title:   "Updated article title",
			Content: "Updated article content",
		}
		requestBody, _ := json.Marshal(articleRequest)

		request, _ := http.NewRequest(http.MethodPut, "/article/"+articleID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return article not found", func(t *testing.T) {
		articleID := "999"
		articleRequest := requests.ArticleRequest{
			Title:   "Updated article title",
			Content: "Updated article content",
		}
		requestBody, _ := json.Marshal(articleRequest)

		request, _ := http.NewRequest(http.MethodPut, "/article/"+articleID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("It should return bad request error", func(t *testing.T) {
		articleID := "1"
		articleRequest := requests.ArticleRequest{}
		requestBody, _ := json.Marshal(articleRequest)

		request, _ := http.NewRequest(http.MethodPut, "/article/"+articleID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestDeleteArticle(t *testing.T) {
	router, articleController := NewArticleTest()

	router.DELETE("/article/:articleID", articleController.Delete)

	t.Run("It should delete one article by articleID", func(t *testing.T) {
		articleID := "1"

		request, _ := http.NewRequest(http.MethodDelete, "/article/"+articleID, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return article not found", func(t *testing.T) {
		articleID := "999"

		request, _ := http.NewRequest(http.MethodDelete, "/article/"+articleID, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
