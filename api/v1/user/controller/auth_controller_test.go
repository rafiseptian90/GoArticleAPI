package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"github.com/rafiseptian90/GoArticle/pkg/config"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewAuthTest() (*gin.Engine, *AuthController) {
	if err := godotenv.Load("../../../../.env"); err != nil {
		log.Fatalf("Cannot load the .env file, error : %v", err.Error())
	}

	db := config.DBConnection()
	router := gin.Default()
	authController := NewAuthController(db)

	return router, authController
}

func TestRegister(t *testing.T) {
	router, authController := NewAuthTest()
	router.POST("/auth/register", authController.Register)

	t.Run("Should return successfully registered", func(t *testing.T) {

		userRequest := map[string]interface{}{
			"username": "rshme",
			"email":    "rshme@me.com",
			"password": "rshme123",
			"name":     "Rafi Septian Hadi",
			"bio":      "Test BIO",
		}

		requestBody, _ := json.Marshal(userRequest)

		request, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("Should return bad request error", func(t *testing.T) {

		userRequest := map[string]interface{}{
			"username": "rshme",
			"password": "rshme123",
			"name":     "Rafi Septian Hadi",
			"bio":      "Test BIO",
		}

		requestBody, _ := json.Marshal(userRequest)

		request, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(requestBody))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
