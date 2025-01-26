package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	go_blog "go-blog"
	"strings"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) CreatePost(post go_blog.Post) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, content, user_id) VALUES ($1, $2, $3) RETURNING id", postsTable)
	row := r.db.QueryRow(query, post.Title, post.Content, post.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *PostPostgres) GetAllPosts() ([]*go_blog.Post, error) {
	var posts []*go_blog.Post
	query := fmt.Sprintf("select p.id as id, p.title as title, p.content as content, p.user_id as user_id, count(l.id) as likes from %s p left join %s l ON p.id = l.post_id group by p.id", postsTable, likesTable)
	if err := r.db.Select(&posts, query); err != nil {
		return make([]*go_blog.Post, 0), nil
	}
	return posts, nil
}

func (r *PostPostgres) GetPost(id int) (go_blog.Post, error) {
	var post go_blog.Post
	query := fmt.Sprintf("SELECT * FROM %s p WHERE p.id = $1", postsTable)
	if err := r.db.Get(&post, query, id); err != nil {
		return post, err
	}
	return post, nil
}

func (r *PostPostgres) DeletePost(id, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s p WHERE p.id = $1 AND p.user_id = $2", postsTable)
	_, err := r.db.Exec(query, id, userId)
	return err
}

func (r *PostPostgres) UpdatePost(id int, userId int, input go_blog.PostUpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s p SET %s WHERE p.id = $%d AND p.user_id = $%d", postsTable, setQuery, argId, argId+1)
	args = append(args, id, userId)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostPostgres) CreateLike(input go_blog.Like) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_id, user_id) VALUES ($1, $2) RETURNING id", likesTable)
	row := r.db.QueryRow(query, input.PostId, input.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostPostgres) DeleteLike(id, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s l WHERE l.id = $1 AND l.user_id = $2", likesTable)
	_, err := r.db.Exec(query, id, userId)
	return err
}
