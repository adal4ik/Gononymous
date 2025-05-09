package driverports

import (
	"context"
	"time"

	"backend/internal/core/domains/dto"
)

type PostDriverPortInterface interface {
	AddPost(post dto.PostDto, data []byte) error
	GetAll() ([]dto.PostDto, error)
	GetPostById(id string) (dto.PostDto, error)
	StartPostArchiver(ctx context.Context, interval time.Duration)
}
