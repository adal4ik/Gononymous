package db

import (
	"backend/internal/core/domains/dao"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (postRepository *PostRepository) AddPost(post dao.PostDao) error {
	sqlQuery := `INSERT INTO posts(post_id, user_id,  title, subject, content, image_url, status)
				 VALUES ($1, $2, $3, $4, $5, $6, $7);`
	_, err := postRepository.db.Exec(sqlQuery, post.PostId, post.UserId, post.Title, post.Subject, post.Content, post.ImageUrl, post.Status)
	if err != nil {
		return err
	}
	return nil
}

func (postRepository *PostRepository) GetActive() ([]dao.PostDao, error) {
	sqlQuery := `SELECT p.post_id, p.created_at, u.name, u.avatar_url, p.title, p.subject, p.content, p.image_url 
				FROM posts as p
				JOIN users as u on u.user_id = p.user_id
				WHERE p.status = 'Active';`
	rows, err := postRepository.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allPosts []dao.PostDao

	for rows.Next() {
		var post dao.PostDao

		err = rows.Scan(&post.PostId, &post.CreatedAt, &post.UserName, &post.UserAvaUrl, &post.Title, &post.Subject, &post.Content, &post.ImageUrl)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}
	return allPosts, nil
}

func (postRepository *PostRepository) GetAll() ([]dao.PostDao, error) {
	sqlQuery := `SELECT p.post_id, p.created_at, u.name, u.avatar_url, p.title, p.subject, p.content, p.image_url 
				FROM posts as p
				JOIN users as u on u.user_id = p.user_id;`
	rows, err := postRepository.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allPosts []dao.PostDao

	for rows.Next() {
		var post dao.PostDao

		err = rows.Scan(&post.PostId, &post.CreatedAt, &post.UserName, &post.UserAvaUrl, &post.Title, &post.Subject, &post.Content, &post.ImageUrl)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}
	return allPosts, nil
}

func (postRepository *PostRepository) GetPostById(id string) (dao.PostDao, error) {
	sqlQuery := `SELECT post_id,user_id, created_at, title, subject, content,image_url FROM posts WHERE post_id = $1;`
	row := postRepository.db.QueryRow(sqlQuery, id)

	var post dao.PostDao
	err := row.Scan(&post.PostId, &post.UserId, &post.CreatedAt, &post.Title, &post.Subject, &post.Content, &post.ImageUrl)
	if err != nil {
		return dao.PostDao{}, err
	}
	return post, nil
}

func (r *PostRepository) ArchiveExpiredPosts(ctx context.Context) error {
	query := `
		UPDATE posts
		SET status = 'Archived'
		WHERE status != 'Archived' AND (
			(NOT EXISTS (
				SELECT 1 FROM comments WHERE comments.post_id = posts.post_id
			) AND created_at <= NOW() - INTERVAL '1 minutes')

			OR

			(EXISTS (
				SELECT 1 FROM comments WHERE comments.post_id = posts.post_id
			) AND created_at <= NOW() - INTERVAL '2 minutes')
		);
	`
	log.Print("hi")
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func (r *PostRepository) GetPostsByUserID(userId string) ([]dao.PostDao, error) {
	sqlQuery := `SELECT p.post_id, p.created_at, u.name, u.avatar_url, p.title, p.subject, p.content, p.image_url 
				FROM posts as p
				JOIN users as u on u.user_id = p.user_id
				WHERE p.user_id = $1
				LIMIT 3;`
	rows, err := r.db.Query(sqlQuery, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allPosts []dao.PostDao

	for rows.Next() {
		var post dao.PostDao

		err = rows.Scan(&post.PostId, &post.CreatedAt, &post.UserName, &post.UserAvaUrl, &post.Title, &post.Subject, &post.Content, &post.ImageUrl)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}
	return allPosts, nil
}
