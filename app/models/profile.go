package models

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	Id        uint           `gorm:"primaryKey;autoIncrement;->" json:"id"`
	UserId    uint           `gorm:"not null" json:"-"`
	Name      string         `gorm:"type:varchar(191);not null" json:"name" binding:"required"`
	Bio       string         `gorm:"type:varchar(191);default:NULL" json:"bio"`
	Photo     string         `gorm:"type:varchar(191);default:NULL" json:"photo"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	User      *User          `json:"user,omitempty"`
}

type ProfileRequest struct {
	Name  string `json:"name" binding:"required"`
	Bio   string `json:"bio" binding:"required"`
	Photo string `json:"photo"`
}
