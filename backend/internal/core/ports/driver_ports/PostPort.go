package driverports

import (
	"backend/internal/core/domains/dto"
	"context"
	"time"
)

type PostDriverPortInterface interface {
	AddPost(post dto.PostDto, data []byte) error
	GetActive() ([]dto.PostDto, error)
	GetAll() ([]dto.PostDto, error)
	GetPostById(id string) (dto.PostDto, error)
	StartPostArchiver(ctx context.Context, interval time.Duration)
	GetPostsByUserID(userId string) ([]dto.PostDto, error)
}
