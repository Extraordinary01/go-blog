package service

import (
	go_blog "go-blog"
	"go-blog/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post go_blog.Post) (int, error) {
	return s.repo.CreatePost(post)
}

func (s *PostService) GetAllPosts() ([]*go_blog.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *PostService) GetPost(id int) (go_blog.Post, error) {
	return s.repo.GetPost(id)
}

func (s *PostService) DeletePost(id, userId int) error {
	return s.repo.DeletePost(id, userId)
}

func (s *PostService) UpdatePost(id, userId int, input go_blog.PostUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdatePost(id, userId, input)
}

func (s *PostService) CreateLike(input go_blog.Like) (int, error) {
	return s.repo.CreateLike(input)
}

func (s *PostService) DeleteLike(id, userId int) error {
	return s.repo.DeleteLike(id, userId)
}
