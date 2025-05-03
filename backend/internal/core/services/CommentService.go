package services

import drivenports "backend/internal/core/ports/driven_ports"

type CommentService struct {
	repo           drivenports.DatabasePortInterface
	imageCollector drivenports.ImageCollectionInterface
}

func NewCommentService(repo drivenports.DatabasePortInterface, imageCollector drivenports.ImageCollectionInterface) *CommentService {
	return &CommentService{repo: repo, imageCollector: imageCollector}
}
