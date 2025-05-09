package driverports

import (
	"context"
	"time"

	"backend/internal/core/domains/dto"
)

type PostDriverPortInterface interface {
	AddPost(post dto.PostDto, data []byte, ctx context.Context) error
	GetActive(ctx context.Context) ([]dto.PostDto, error)
	GetAll(ctx context.Context) ([]dto.PostDto, error)
	GetPostById(id string, ctx context.Context) (dto.PostDto, error)
	StartPostArchiver(ctx context.Context, interval time.Duration)
}
