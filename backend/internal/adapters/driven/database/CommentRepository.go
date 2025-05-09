package db

import (
	"backend/internal/core/domains/dto"
	"database/sql"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (commentRepository *CommentRepository) AddComment(comment dto.Comment) error {
	sqlQuery := `INSERT INTO comments(comment_id, post_id, parent_id, user_id, content, image_url)
				VALUES($1, $2, $3, $4, $5, $6);`
	_, err := commentRepository.db.Exec(sqlQuery, comment.CommentID, comment.PostID, comment.ParentID, comment.UserID, comment.Content, comment.ImageUrl)
	if err != nil {
		return err
	}
	return nil
}

func (commentRepository *CommentRepository) GetCommentsByPostId(id string) ([]dto.Comment, error) {
	sqlQuery := `SELECT c.comment_id, c.post_id, c.parent_id, c.user_id, u.name , u.avatar_url , c.content, c.image_url 
				FROM comments as c
				JOIN users as u on u.user_id = c.user_id
				WHERE c.post_id = $1 AND c.parent_id = '00000000-0000-0000-0000-000000000000';`
	rows, err := commentRepository.db.Query(sqlQuery, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []dto.Comment

	for rows.Next() {
		var comment dto.Comment

		err = rows.Scan(&comment.CommentID, &comment.PostID, &comment.ParentID, &comment.UserID, &comment.UserName, &comment.UserAvatarLink, &comment.Content, &comment.ImageUrl)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (commentRepository *CommentRepository) GetCommentReplies(commentId string) ([]dto.Comment, error) {
	sqlQuery := `SELECT c.comment_id, c.post_id, c.parent_id, c.user_id, u.name , u.avatar_url , c.content, c.image_url 
				FROM comments as c
				JOIN users as u on u.user_id = c.user_id
				WHERE c.parent_id = $1;`
	rows, err := commentRepository.db.Query(sqlQuery, commentId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []dto.Comment

	for rows.Next() {
		var comment dto.Comment

		err = rows.Scan(&comment.CommentID, &comment.PostID, &comment.ParentID, &comment.UserID, &comment.UserName, &comment.UserAvatarLink, &comment.Content, &comment.ImageUrl)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
