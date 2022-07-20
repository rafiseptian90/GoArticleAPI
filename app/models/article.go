package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Id        uint           `gorm:"primaryKey;autoIncrement;->" json:"id"`
	Title     string         `gorm:"type:varchar(191);not null" json:"title" binding:"required"`
	Content   string         `gorm:"not null" json:"content" binding:"required"`
	HeroImg   string         `gorm:"type:varchar(191);default:NULL" json:"hero_img"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Tags      []*Tag         `gorm:"many2many:article_tags" json:"tags"`
}
