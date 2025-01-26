package service

import (
	go_blog "go-blog"
	"go-blog/pkg/repository"
)

type Authorization interface {
	CreateUser(user go_blog.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	Post
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Post:          NewPostService(repo.Post),
	}
}
