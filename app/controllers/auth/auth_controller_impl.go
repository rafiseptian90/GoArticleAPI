package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rafiseptian90/GoArticle/app/models"
	"github.com/rafiseptian90/GoArticle/config"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type Controller struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) *Controller {
	return &Controller{
		DB: DB,
	}
}

func (controller *Controller) Login(ctx *gin.Context) {
	type Credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user models.User
	var credentials Credentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	if result := controller.DB.Where("username = ?", credentials.Username).Preload("Profile").Find(&user); result.RowsAffected < 1 {
		ResponseJSON.Unauthorized(ctx, "Username is not found !")
		return
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		ResponseJSON.Unauthorized(ctx, "Password is not match !")
		return
	}

	// Remove password element from user
	user.Password = ""

	// Generate the JWT token from username
	jwtToken, err := config.JWTGenerateToken(user.Username)
	if err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}

	// Set the response object
	response := map[string]any{
		"user":  user,
		"token": jwtToken,
	}

	ResponseJSON.SuccessWithData(ctx, "Login successful", response)
}

func (controller *Controller) Register(ctx *gin.Context) {
	var user models.User
	var profile models.Profile

	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := ctx.ShouldBindBodyWith(&profile, binding.JSON); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	if err := controller.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&user); result.RowsAffected < 1 {
			return result.Error
		}

		profile.UserId = user.Id

		if result := tx.Create(&profile); result.RowsAffected < 1 {
			return result.Error
		}

		return nil
	}); err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}
}

func (controller *Controller) ForgotPassword(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *Controller) Refresh(ctx *gin.Context) {
	token := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
	newToken, err := config.JWTRefreshToken("username", token)
	if err != nil {
		ResponseJSON.Unauthorized(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Token has been refreshed", newToken)
}
