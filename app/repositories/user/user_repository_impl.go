package user

import (
	"github.com/rafiseptian90/GoArticle/app/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *Repository {
	return &Repository{DB: DB}
}

func (repository *Repository) GetUserByUsername(username string) (models.User, error) {
	var user models.User

	if query := repository.DB.Preload("Profile").Preload("Articles.User.Profile").Preload("Articles.Tags").Where("username = ?", username).Find(&user); query.RowsAffected < 1 {
		return user, query.Error
	}

	return user, nil
}
