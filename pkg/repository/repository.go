package repository

import (
	"github.com/jmoiron/sqlx"
	go_blog "go-blog"
	"go-blog/pkg/repository/postgres"
)

type Authorization interface {
	CreateUser(user go_blog.User) (int, error)
	GetUser(username, passwordHash string) (go_blog.User, error)
}

type Post interface {
	CreatePost(post go_blog.Post) (int, error)
	GetAllPosts() ([]*go_blog.Post, error)
	GetPost(id int) (go_blog.Post, error)
	DeletePost(id, userId int) error
	UpdatePost(id, userId int, input go_blog.PostUpdateInput) error
	CreateLike(input go_blog.Like) (int, error)
	DeleteLike(id, userId int) error
}

type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		Post:          postgres.NewPostPostgres(db),
	}
}
