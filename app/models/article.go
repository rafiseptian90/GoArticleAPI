package models

import "gorm.io/gorm"

type Article struct {
	Title   string `gorm:"max:191; not null"`
	Content string `gorm:"type:text; not null"`
	HeroImg string
	gorm.Model
}
