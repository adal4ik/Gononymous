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
	sqlQuery := `INSERT INTO posts(post_id, created_at, title, subject, content, image_url, status)
				 VALUES ($1, $2, $3, $4, $5, $6, $7);`
	_, err := postRepository.db.Exec(sqlQuery, post.PostId, post.CreatedAt, post.Title, post.Subject, post.Content, post.ImageUrl, post.Status)
	if err != nil {
		return err
	}
	return nil
}

func (postRepository *PostRepository) GetAll() ([]dao.PostDao, error) {
	sqlQuery := `SELECT post_id, created_at, title, subject, content FROM posts;`
	rows, err := postRepository.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allPosts []dao.PostDao

	for rows.Next() {
		var post dao.PostDao

		err = rows.Scan(&post.PostId, &post.CreatedAt, &post.Title, &post.Subject, &post.Content)

		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}
	return allPosts, nil
}
