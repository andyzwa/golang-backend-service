package usecases

import (
	"golang-backend-service/src/domain"
)

type articleUseCase struct {
	articleRepo domain.ArticleRepository
}

// NewArticleUseCase will create new an articleUseCase object representation of domain.ArticleUseCase interface
func NewArticleUseCase(a domain.ArticleUseCase) domain.ArticleUseCase {
	return &articleUseCase{
		articleRepo: a,
	}
}

func (a articleUseCase) FindAll() (domain.Articles, error) {
	return a.articleRepo.FindAll()
}

func (a articleUseCase) FindByID(id string) (domain.Article, error) {
	return a.articleRepo.FindByID(id)
}

func (a articleUseCase) DeleteByID(id string) error {
	return a.articleRepo.DeleteByID(id)
}

func (a articleUseCase) Create(article domain.Article) (domain.Article, error) {
	return a.articleRepo.Create(article)
}
