package fakers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/rafiseptian90/GoArticle/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewUserSeeders(DB *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	for i := 1; i <= 5; i++ {
		var user = &models.User{
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: string(password),
		}
		DB.Create(user)

		var profile = models.Profile{
			UserId: user.Id,
			Name:   faker.Name(),
		}
		DB.Create(&profile)
	}
}
