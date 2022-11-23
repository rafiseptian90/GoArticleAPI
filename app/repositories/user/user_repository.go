package user

import "github.com/rafiseptian90/GoArticle/app/models"

type UserRepository interface {
	GetUserByUsername(username string) (models.User, error)
}
