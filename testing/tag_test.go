package testing

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafiseptian90/GoArticle/app/controllers/tag"
	"github.com/rafiseptian90/GoArticle/app/models"
	tag2 "github.com/rafiseptian90/GoArticle/app/repositories/tag"
	"github.com/rafiseptian90/GoArticle/config"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTagTest() (*gin.Engine, *tag.Controller) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := config.DBConnection()
	router := gin.Default()
	tagRepository := tag2.NewTagRepository(DB)
	tagController := tag.NewTagController(tagRepository)

	return router, tagController
}

func TestGetTags(t *testing.T) {
	router, tagController := NewTagTest()

	router.GET("/tag", tagController.Index)
	request, _ := http.NewRequest(http.MethodGet, "/tag", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//response, _ := io.ReadAll(recorder.Body)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestStoreTag(t *testing.T) {
	router, tagController := NewTagTest()
	router.POST("/tag", tagController.Store)

	t.Run("It should create a new tag", func(t *testing.T) {
		tagRequest := models.Tag{
			Name: "New Tag",
		}
		requestBody, _ := json.Marshal(tagRequest)

		request, _ := http.NewRequest(http.MethodPost, "/tag", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return a bad request error", func(t *testing.T) {
		tagRequest := models.Tag{}
		requestBody, _ := json.Marshal(tagRequest)

		request, _ := http.NewRequest(http.MethodPost, "/tag", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetTag(t *testing.T) {
	router, tagController := NewTagTest()

	router.GET("/tag/:tagID", tagController.Show)

	t.Run("It should find a tag", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tag/1", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It shouldn't find a tag", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tag/9999", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func TestUpdateTag(t *testing.T) {
	router, tagController := NewTagTest()
	router.PUT("/tag/:tagID", tagController.Update)

	t.Run("It should update a tag", func(t *testing.T) {
		tagRequest := models.Tag{
			Name: "Updated Tag",
		}
		tagID := "1"
		requestBody, _ := json.Marshal(tagRequest)

		request, _ := http.NewRequest(http.MethodPut, "/tag/"+tagID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return a tag not found", func(t *testing.T) {
		tagRequest := models.Tag{
			Name: "Updated Tag",
		}
		tagID := "100"
		requestBody, _ := json.Marshal(tagRequest)

		request, _ := http.NewRequest(http.MethodPut, "/tag/"+tagID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("It should return a bad request error", func(t *testing.T) {
		tagRequest := models.Tag{}
		tagID := "1"
		requestBody, _ := json.Marshal(tagRequest)

		request, _ := http.NewRequest(http.MethodPut, "/tag/"+tagID, bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestDeleteTag(t *testing.T) {
	router, tagController := NewTagTest()

	router.DELETE("/tag/:tagID", tagController.Delete)

	t.Run("It should delete the tag", func(t *testing.T) {
		tagID := "1"

		request, _ := http.NewRequest(http.MethodDelete, "/tag/"+tagID, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("It should return tag not found", func(t *testing.T) {
		tagID := "100"

		request, _ := http.NewRequest(http.MethodDelete, "/tag/"+tagID, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
