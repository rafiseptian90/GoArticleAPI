package testing

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafiseptian90/GoArticle/app/controllers/auth"
	"github.com/rafiseptian90/GoArticle/config"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewAuthTest() (*gin.Engine, *auth.Controller) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Cannot load the .env file, error : %v", err.Error())
	}

	db := config.DBConnection()
	router := gin.Default()
	authController := auth.NewAuthController(db)

	return router, authController
}

func TestRegister(t *testing.T) {
	router, authController := NewAuthTest()

	router.POST("/auth/register", authController.Register)

	request, _ := http.NewRequest(http.MethodPost, "/auth/register", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
