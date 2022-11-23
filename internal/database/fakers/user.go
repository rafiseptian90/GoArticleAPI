package fakers

import (
	"github.com/bxcodec/faker/v3"
	models2 "github.com/rafiseptian90/GoArticle/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewUserSeeders(DB *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	for i := 1; i <= 5; i++ {
		var user = &models2.User{
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: string(password),
		}
		DB.Create(user)

		var profile = models2.Profile{
			UserId: user.Id,
			Name:   faker.Name(),
		}
		DB.Create(&profile)
	}
}
