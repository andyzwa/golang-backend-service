package repository

import (
	"fmt"
	"golang-backend-service/src/domain"
	"strconv"
)

// inMemoryArticlesRepository implements domain.ArticleRepository
type inMemoryArticlesRepository struct {
	articles domain.Articles
}

func NewInMemoryArticlesRepository() domain.ArticleRepository {
	return &inMemoryArticlesRepository{

		//Init some example articles
		articles: []domain.Article{
			{Id: "1", Title: "Article 1", Content: "Article Content 1", Desc: "#ff2"},
			{Id: "2", Title: "Article 2", Content: "Article Content 2", Desc: "#0bdcab"},
		}}
}

func (ir *inMemoryArticlesRepository) FindAll() (domain.Articles, error) {
	return ir.articles, nil
}

func (ir *inMemoryArticlesRepository) FindByID(id string) (domain.Article, error) {
	for _, article := range ir.articles {
		if article.Id == id {
			return article, nil
		}
	}
	return domain.Article{}, fmt.Errorf("no article found")
}

func (ir *inMemoryArticlesRepository) DeleteByID(id string) error {
	for index, article := range ir.articles {
		if article.Id == id {
			ir.articles = append(ir.articles[:index], ir.articles[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no article found to delete")
}

func (ir *inMemoryArticlesRepository) Create(article domain.Article) (domain.Article, error) {

	//Append to articles if new
	if article.Id == "" {
		article.Id = strconv.Itoa(len(ir.articles) + 1)
		ir.articles = append(ir.articles, article)
	} else {
		//update
		for index, art := range ir.articles {
			if art.Id == article.Id {
				ir.articles[index].Desc = article.Desc
				ir.articles[index].Title = article.Title
				ir.articles[index].Content = article.Content
			}
		}
	}
	return article, nil
}
