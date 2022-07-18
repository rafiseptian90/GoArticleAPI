package requests

type Tabler interface {
	TableName() string
}

type CategoryRequest struct {
	Name string `form:"name" binding:"required"`
}

// TableName overrides the table name used by CategoryResponse to `categories`
func (CategoryRequest) TableName() string {
	return "categories"
}
