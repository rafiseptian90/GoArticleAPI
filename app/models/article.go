package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Id        uint           `gorm:"primaryKey;autoIncrement;->" json:"id"`
	UserId    uint           `gorm:"not null" json:"user_id"`
	Title     string         `gorm:"type:varchar(191);not null" json:"title" binding:"required"`
	Content   string         `gorm:"not null" json:"content" binding:"required"`
	Thumbnail string         `gorm:"type:varchar(191);default:NULL" json:"thumbnail"`
	Seen      uint           `gorm:"default:0" json:"seen"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	User      *User          `json:"user,omitempty"`
	Tags      []Tag          `gorm:"many2many:article_tags" json:"tags"`
}
