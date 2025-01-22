package models

import "github.com/go-pg/pg/v10"

type Author struct {
	ID int64 `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
}

func (a *Author) String() string {
	return a.FirstName + " " + a.LastName
}

func (a *Author) Create(db *pg.DB) error {
	_, err := db.Model(a).Insert()
	return err
}

func (a *Author) Get(db *pg.DB, authorId int64) error {
	return db.Model(a).Where("id = ?", authorId).Select()
}

func (a *Author) Update(db *pg.DB, authorId int64) error {
	_, err := db.Model(a).Where("id = ?", authorId).Update()
	return err
}

func (a *Author) Delete(db *pg.DB, authorId int64) error {
	_, err := db.Model(a).Where("id = ?", authorId).Delete()
	return err
}

func GetAuthors(db *pg.DB) ([]*Author, error) {
	authors := make([]*Author, 0)
	err := db.Model(&authors).Select()
	return authors, err
}