package go_blog

import "errors"

type Post struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" binding:"required" db:"title"`
	Content string `json:"content" binding:"required" db:"content"`
	UserId  int    `json:"user_id" db:"user_id"`
	Likes   int    `json:"likes,omitempty" db:"likes"`
}

type Like struct {
	Id     int `json:"-" db:"id"`
	UserId int `json:"user_id"`
	PostId int `json:"post_id" binding:"required"`
}

type PostUpdateInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (i PostUpdateInput) Validate() error {
	if i.Title == nil && i.Content == nil {
		return errors.New("no data provided")
	}
	return nil
}
