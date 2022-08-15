package models

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/config"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `gorm:"primaryKey;autoIncrement;->" json:"id"`
	Username  string         `gorm:"type:varchar(191);not null;unique" json:"username" binding:"required"`
	Email     string         `gorm:"varchar(191);not null;unique" json:"email" binding:"required"`
	Password  string         `gorm:"type:varchar(191);not null" json:"password,omitempty" binding:"required"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Profile   *Profile       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"profile"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func AuthUser(ctx *gin.Context) *User {
	var user User
	DB := config.DBConnection()

	if result := DB.Where("email = ?", ctx.Value("user.email")).Find(&user); result.RowsAffected < 1 {
		ResponseJSON.InternalServerError(ctx, result.Error.Error())
		return nil
	}

	return &user
}
