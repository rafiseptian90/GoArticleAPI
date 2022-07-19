package requests

type ArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	HeroImg string `json:"hero_img"`
}

// TableName overrides the table name used by ArticleRequest to `articles`
func (ArticleRequest) TableName() string {
	return "articles"
}
