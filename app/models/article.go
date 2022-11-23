package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Id          uint           `gorm:"primaryKey;autoIncrement;->" json:"id"`
	UserId      uint           `gorm:"not null" json:"-"`
	Slug        string         `gorm:"type:varchar(191);not null" json:"slug"`
	Title       string         `gorm:"type:varchar(191);not null" json:"title"`
	Content     string         `gorm:"not null" json:"content"`
	Thumbnail   string         `gorm:"type:varchar(191);default:NULL" json:"thumbnail"`
	Seen        uint           `gorm:"default:0" json:"seen"`
	PublishedAt time.Time      `gorm:"nullable;default:NULL" json:"published_at"`
	CreatedAt   *time.Time     `gorm:"default:now()" json:"-"`
	UpdatedAt   *time.Time     `gorm:"default:now()" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	User        *User          `json:"creator,omitempty"`
	Tags        []Tag          `gorm:"many2many:article_tags" json:"tags"`
}

type ArticleRequest struct {
	UserId    uint   `json:"user_id"`
	Slug      string `json:"slug"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Thumbnail string `json:"thumbnail"`
}
