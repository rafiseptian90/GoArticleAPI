package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Id        uint           `gorm:"primaryKey;autoIncrement;->" json:"id"`
	Slug      string         `gorm:"type:varchar(191);not null" json:"slug"`
	Name      string         `gorm:"type:varchar(191);not null" json:"name"`
	CreatedAt time.Time      `gorm:"default:now()" json:"-"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Articles  []Article      `gorm:"many2many:article_tags" json:"articles"`
}

type TagRequest struct {
	Slug string `json:"slug"`
	Name string `json:"name" binding:"required"`
}

type TagRepository interface {
	GetTags() []Tag
	GetTag(slug string) (Tag, error)
	StoreTag(tagRequest *TagRequest) error
	UpdateTag(slug string, tagRequest *TagRequest) error
	DeleteTag(slug string) error
}
