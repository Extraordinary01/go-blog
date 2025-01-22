package models

import "github.com/go-pg/pg/v10"

type Blog struct {
	ID      int64  `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	AuthorId int64 `json:"author_id" binding:"required"`
	Author  *Author `pg:"rel:has-one" json:"author,omitempty"`
}

func (b *Blog) Create(db *pg.DB) error {
	_, err := db.Model(b).Insert()
	return err
}

func (b *Blog) Get(db *pg.DB, blogId int64) error {
	return db.Model(b).Where("blog.id = ?", blogId).Relation("Author").Select()
}

func (b *Blog) Update(db *pg.DB, blogId int64) error {
	_, err := db.Model(b).Where("blog.id = ?", blogId).Update()
	return err
}

func (b *Blog) Delete(db *pg.DB, blogId int64) error {
	_, err := db.Model(b).Where("blog.id = ?", blogId).Delete()
	return err
}

func GetBlogs(db *pg.DB) ([]*Blog, error) {
	blogs := make([]*Blog, 0)
	err := db.Model(&blogs).Relation("Author").Select()
	return blogs, err
}