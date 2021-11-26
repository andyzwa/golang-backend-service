package usecases

import (
	"golang-backend-service/src/domain"
	"golang-backend-service/src/repository"
	"reflect"
	"testing"
)

func TestNewArticleUseCase(t *testing.T) {
	type args struct {
		a domain.ArticleUseCase
	}
	tests := []struct {
		name string
		args args
		want domain.ArticleUseCase
	}{
		{name: "Test1", args: args{NewArticleUseCase(nil)}, want: NewArticleUseCase(&articleUseCase{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArticleUseCase(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticleUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleUseCase_Create(t *testing.T) {
	type fields struct {
		articleRepo domain.ArticleRepository
	}
	type args struct {
		article domain.Article
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Article
		wantErr bool
	}{
		{name: "Test1",
			fields: struct{ articleRepo domain.ArticleRepository }{articleRepo: repository.NewInMemoryArticlesRepository()},
			args: struct{ article domain.Article }{article: domain.Article{
				Id:      "3",
				Title:   "Art3",
				Content: "Art3",
				Desc:    "Art3",
			}},
			want:    domain.Article{Id: "3", Title: "Art3", Content: "Art3", Desc: "Art3"},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := articleUseCase{
				articleRepo: tt.fields.articleRepo,
			}
			got, err := a.Create(tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleUseCase_DeleteByID(t *testing.T) {
	type fields struct {
		articleRepo domain.ArticleRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Test1",
			fields:  struct{ articleRepo domain.ArticleRepository }{articleRepo: repository.NewInMemoryArticlesRepository()},
			args:    struct{ id string }{id: "1"},
			wantErr: false},
		{name: "Test2",
			fields:  struct{ articleRepo domain.ArticleRepository }{articleRepo: repository.NewInMemoryArticlesRepository()},
			args:    struct{ id string }{id: "3"},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := articleUseCase{
				articleRepo: tt.fields.articleRepo,
			}
			if err := a.DeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_articleUseCase_FindAll(t *testing.T) {
	type fields struct {
		articleRepo domain.ArticleRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.Articles
		wantErr bool
	}{
		{name: "Test1",
			fields: struct{ articleRepo domain.ArticleRepository }{articleRepo: repository.NewInMemoryArticlesRepository()},
			want: domain.Articles{
				domain.Article{Id: "1", Title: "Article 1", Content: "Article Content 1", Desc: "#ff2"},
				domain.Article{Id: "2", Title: "Article 2", Content: "Article Content 2", Desc: "#0bdcab"},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := articleUseCase{
				articleRepo: tt.fields.articleRepo,
			}
			got, err := a.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleUseCase_FindByID(t *testing.T) {
	type fields struct {
		articleRepo domain.ArticleRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Article
		wantErr bool
	}{
		{name: "Test1",
			fields:  struct{ articleRepo domain.ArticleRepository }{articleRepo: repository.NewInMemoryArticlesRepository()},
			args:    struct{ id string }{id: "1"},
			want:    domain.Article{Id: "1", Title: "Article 1", Content: "Article Content 1", Desc: "#ff2"},
			wantErr: false},
		{name: "Test2",
			fields:  struct{ articleRepo domain.ArticleRepository }{articleRepo: repository.NewInMemoryArticlesRepository()},
			args:    struct{ id string }{id: "3"},
			want:    domain.Article{},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := articleUseCase{
				articleRepo: tt.fields.articleRepo,
			}
			got, err := a.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
