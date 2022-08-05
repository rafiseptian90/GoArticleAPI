package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Id        uint           `gorm:"primaryKey;autoIncrement;->" json:"id"`
	Name      string         `gorm:"type:varchar(191);not null" json:"name" binding:"required"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
