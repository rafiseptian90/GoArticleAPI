package models

import "gorm.io/gorm"

type Category struct {
	Name string
	gorm.Model
}
