package domain

// Article - Our struct for all articles
// swagger:model Article
type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Desc    string `json:"color"`
}

// swagger:model Articles
type Articles []Article

// ArticleRepository definition
type ArticleRepository interface {
	FindAll() (Articles, error)
	FindByID(string) (Article, error)
	DeleteByID(string) error
	Create(Article) (Article, error)
}

// ArticleUseCase definition
type ArticleUseCase interface {
	FindAll() (Articles, error)
	FindByID(string) (Article, error)
	DeleteByID(string) error
	Create(Article) (Article, error)
}
