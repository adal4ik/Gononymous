package drivenports

import (
	"context"

	"backend/internal/core/domains/dto"
)

type CommentRepoInterface interface {
	AddComment(comment dto.Comment, ctx context.Context) error
	GetCommentsByPostId(id string, ctx context.Context) ([]dto.Comment, error)
	GetCommentReplies(commentId string, ctx context.Context) ([]dto.Comment, error)
}
