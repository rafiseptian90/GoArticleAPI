package responses

import "gorm.io/gorm"

type ArticleResponse struct {
	Id        int            `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	HeroImg   string         `json:"hero_img"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// TableName overrides the table name used by ArticleResponse to `articles`
func (ArticleResponse) TableName() string {
	return "articles"
}
