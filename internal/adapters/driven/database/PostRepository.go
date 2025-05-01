package db

import (
	"Gononymous/internal/core/domains/dao"
	"database/sql"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (postRepository *PostRepository) AddPost(post dao.PostDao) error {
	sqlQuery := `INSERT INTO posts(post_id, created_at, title, subject, content, image_url)
				 VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := postRepository.db.Exec(sqlQuery, post.PostId, post.CreatedAt, post.Title, post.Subject, post.Content, post.ImageUrl)
	if err != nil {
		return err
	}
	return nil
}
