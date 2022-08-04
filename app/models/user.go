package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `gorm:"primaryKey;autoIncrement;->" json:"id"`
	Username  string         `gorm:"type:varchar(191);not null;unique" json:"username" binding:"required"`
	Password  string         `gorm:"type:varchar(191);not null;<-" json:"password" binding:"required"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Profile   *Profile       `gorm:"constraint:OnUpdate:CASCADE:OnDelete:SET NULL;"`
}
