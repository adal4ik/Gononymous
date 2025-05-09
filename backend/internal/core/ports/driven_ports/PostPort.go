package drivenports

import (
	"context"

	"backend/internal/core/domains/dao"
)

type DatabasePortInterface interface {
	AddPost(post dao.PostDao, ctx context.Context) error
	GetActive(ctx context.Context) ([]dao.PostDao, error)
	GetAll(ctx context.Context) ([]dao.PostDao, error)
	GetPostById(id string, ctx context.Context) (dao.PostDao, error)
	ArchiveExpiredPosts(ctx context.Context) error
}
