package repository

import (
	"errors"
	"github.com/rafiseptian90/GoArticle/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (repository UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	if query := repository.DB.Preload("Profile").Preload("Articles.User.Profile").Preload("Articles.Tags").Where("username = ?", username).Find(&user); query.RowsAffected < 1 {
		return &user, errors.New("username isn't found")
	}

	return &user, nil
}
