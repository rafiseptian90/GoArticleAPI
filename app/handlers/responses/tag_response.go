package responses

import "gorm.io/gorm"

type Tabler interface {
	TableName() string
}

type TagResponse struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// TableName overrides the table name used by TagResponse to `categories`
func (TagResponse) TableName() string {
	return "categories"
}
