package driverports

import "backend/internal/core/domains/dto"

type CommentServiceInterface interface {
	AddComment(comment dto.Comment) error
}
