package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rafiseptian90/GoArticle/app/models"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	//TODO implement me
	panic("implement me")
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

func (controller *Controller) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
