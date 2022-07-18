package responses

import "gorm.io/gorm"

type Tabler interface {
	TableName() string
}

type CategoryResponse struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// TableName overrides the table name used by CategoryResponse to `categories`
func (CategoryResponse) TableName() string {
	return "categories"
}
