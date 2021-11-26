package repository

import (
	"golang-backend-service/src/domain"
	"reflect"
	"testing"
)

func TestNewInMemoryArticlesRepository(t *testing.T) {
	tests := []struct {
		name string
		want domain.ArticleRepository
	}{
		{name: "Test1", want: NewInMemoryArticlesRepository()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemoryArticlesRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemoryArticlesRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemoryArticlesRepository_Create(t *testing.T) {
	type fields struct {
		articles domain.Articles
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
			fields: struct{ articles domain.Articles }{articles: []domain.Article{{Id: "1", Title: "Art1", Content: "Art1", Desc: "Art1"}}},
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
			ir := &inMemoryArticlesRepository{
				articles: tt.fields.articles,
			}
			got, err := ir.Create(tt.args.article)
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

func Test_inMemoryArticlesRepository_DeleteByID(t *testing.T) {
	type fields struct {
		articles domain.Articles
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
		{
			name:    "Test1",
			fields:  struct{ articles domain.Articles }{articles: []domain.Article{{Id: "1", Title: "Art1", Content: "Art1", Desc: "Art1"}}},
			args:    args{"1"},
			wantErr: false,
		},
		{
			name:    "Test2",
			fields:  struct{ articles domain.Articles }{articles: []domain.Article{{Id: "1", Title: "Art1", Content: "Art1", Desc: "Art1"}}},
			args:    args{"2"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ir := &inMemoryArticlesRepository{
				articles: tt.fields.articles,
			}
			if err := ir.DeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inMemoryArticlesRepository_FindAll(t *testing.T) {
	type fields struct {
		articles domain.Articles
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.Articles
		wantErr bool
	}{
		{name: "Test1", fields: struct{ articles domain.Articles }{articles: []domain.Article{{Id: "1", Title: "Art1", Content: "Art1", Desc: "Art1"}}},
			want:    []domain.Article{{Id: "1", Title: "Art1", Content: "Art1", Desc: "Art1"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ir := &inMemoryArticlesRepository{
				articles: tt.fields.articles,
			}
			got, err := ir.FindAll()
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

func Test_inMemoryArticlesRepository_FindByID(t *testing.T) {
	type fields struct {
		articles domain.Articles
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
		{
			name:    "Test1",
			fields:  struct{ articles domain.Articles }{articles: []domain.Article{{Id: "1", Title: "Art1", Content: "Art1", Desc: "Art1"}}},
			args:    args{"1"},
			want:    domain.Article{Id: "1", Title: "Art1", Content: "Art1", Desc: "Art1"},
			wantErr: false,
		},
		{
			name:    "Test2",
			fields:  struct{ articles domain.Articles }{articles: []domain.Article{{Id: "1", Title: "Art1", Content: "Art1", Desc: "Art1"}}},
			args:    args{"2"},
			want:    domain.Article{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ir := &inMemoryArticlesRepository{
				articles: tt.fields.articles,
			}
			got, err := ir.FindByID(tt.args.id)
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
