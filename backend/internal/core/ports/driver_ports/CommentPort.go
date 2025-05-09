package driverports

import "backend/internal/core/domains/dto"

type CommentServiceInterface interface {
	AddComment(comment dto.Comment, img []byte) error
	GetCommentsByPostId(postId string) ([]dto.Comment, error)
}
