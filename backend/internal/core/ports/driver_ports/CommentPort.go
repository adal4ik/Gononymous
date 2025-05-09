package driverports

import (
	"context"

	"backend/internal/core/domains/dto"
)

type CommentServiceInterface interface {
	AddComment(comment dto.Comment, img []byte, ctx context.Context) error
	GetCommentsByPostId(postId string,ctx context.Context) ([]dto.Comment, error)
}
