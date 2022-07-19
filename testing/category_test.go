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

func NewCategoryTest() (*gin.Engine, *controllers.CategoryController) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := config.DBConnection()
	router := gin.Default()
	categoryRepository := repositories.NewCategoryRepository(DB)
	categoryController := controllers.NewCategoryController(categoryRepository)

	return router, categoryController
}

func TestGetCategories(t *testing.T) {
	router, categoryController := NewCategoryTest()

	router.GET("/category", categoryController.Index)
	request, _ := http.NewRequest(http.MethodGet, "/category", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//response, _ := io.ReadAll(recorder.Body)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetCategory(t *testing.T) {
	router, categoryController := NewCategoryTest()

	router.GET("/category/:categoryID", categoryController.Show)

	t.Run("It should find a category", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/category/1", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It shouldn't find a category", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/category/9999", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func TestStoreCategory(t *testing.T) {
	router, categoryController := NewCategoryTest()
	router.POST("/category", categoryController.Store)

	t.Run("It should create a new category", func(t *testing.T) {
		categoryRequest := requests.CategoryRequest{
			Name: "New Category",
		}
		requestBody, _ := json.Marshal(categoryRequest)

		request, _ := http.NewRequest(http.MethodPost, "/category", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return a bad request error", func(t *testing.T) {
		categoryRequest := requests.CategoryRequest{}
		requestBody, _ := json.Marshal(categoryRequest)

		request, _ := http.NewRequest(http.MethodPost, "/category", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestUpdateCategory(t *testing.T) {
	router, categoryController := NewCategoryTest()
	router.PUT("/category/:categoryID", categoryController.Update)

	t.Run("It should update a category", func(t *testing.T) {
		categoryRequest := requests.CategoryRequest{
			Name: "Updated Category",
		}
		categoryID := "1"
		requestBody, _ := json.Marshal(categoryRequest)

		request, _ := http.NewRequest(http.MethodPut, "/category/"+categoryID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return a category not found", func(t *testing.T) {
		categoryRequest := requests.CategoryRequest{
			Name: "Updated Category",
		}
		categoryID := "100"
		requestBody, _ := json.Marshal(categoryRequest)

		request, _ := http.NewRequest(http.MethodPut, "/category/"+categoryID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("It should return a bad request error", func(t *testing.T) {
		categoryRequest := requests.CategoryRequest{}
		categoryID := "1"
		requestBody, _ := json.Marshal(categoryRequest)

		request, _ := http.NewRequest(http.MethodPut, "/category/"+categoryID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestDeleteCategory(t *testing.T) {
	router, categoryController := NewCategoryTest()

	router.DELETE("/category/:categoryID", categoryController.Delete)

	t.Run("It should delete the category", func(t *testing.T) {
		categoryID := "1"

		request, _ := http.NewRequest(http.MethodDelete, "/category/"+categoryID, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return category not found", func(t *testing.T) {
		categoryID := "100"

		request, _ := http.NewRequest(http.MethodDelete, "/category/"+categoryID, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
