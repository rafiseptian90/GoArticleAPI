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
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user models.User
	var credentials Credentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	if result := controller.DB.Where("email = ?", credentials.Email).Preload("Profile").Find(&user); result.RowsAffected < 1 {
		ResponseJSON.Unauthorized(ctx, "Email is not found !")
		return
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		ResponseJSON.Unauthorized(ctx, "Password is not match !")
		return
	}

	// Remove password element from user
	user.Password = ""

	// Generate the JWT token from email
	jwtToken, err := config.JWTGenerateToken(user.Email)
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

func (controller *Controller) UpdateProfile(ctx *gin.Context) {
	authUser := models.AuthUser(ctx)
	var userRequest models.UserRequest
	var profileRequest models.ProfileRequest
	var user models.User
	var profile models.Profile

	if err := ctx.ShouldBindBodyWith(&userRequest, binding.JSON); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}
	if err := ctx.ShouldBindBodyWith(&profileRequest, binding.JSON); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	// Update user
	if result := controller.DB.Model(&user).Where("email = ?", authUser.Email).Updates(map[string]interface{}{
		"username": userRequest.Username,
		"email":    userRequest.Email,
	}); result.RowsAffected < 1 {
		ResponseJSON.InternalServerError(ctx, result.Error.Error())
		return
	}

	// Update profile
	if result := controller.DB.Model(&profile).Where("user_id = ?", authUser.Id).Updates(map[string]interface{}{
		"name": profileRequest.Name,
		"bio":  profileRequest.Bio,
	}); result.RowsAffected < 1 {
		ResponseJSON.InternalServerError(ctx, result.Error.Error())
		return
	}

	ResponseJSON.Success(ctx, "Profile has been updated")
	return
}

func (controller *Controller) ForgotPassword(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *Controller) Refresh(ctx *gin.Context) {
	token := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
	newToken, err := config.JWTRefreshToken("email", token)
	if err != nil {
		ResponseJSON.Unauthorized(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Token has been refreshed", newToken)
}
