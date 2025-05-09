package drivenports

import "backend/internal/core/domains/dto"

type CommentRepoInterface interface {
	AddComment(comment dto.Comment) error
	GetCommentsByPostId(id string) ([]dto.Comment, error)
	GetCommentReplies(commentId string) ([]dto.Comment, error)
}
