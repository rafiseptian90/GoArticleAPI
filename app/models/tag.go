package models

import "gorm.io/gorm"

type Tag struct {
	Name string
	gorm.Model
}
