package drivenports

import (
	"backend/internal/core/domains/dao"
	"context"
)

type DatabasePortInterface interface {
	AddPost(post dao.PostDao, ctx context.Context) error
	GetActive(ctx context.Context) ([]dao.PostDao, error)
	GetAll(ctx context.Context) ([]dao.PostDao, error)
	GetPostById(id string, ctx context.Context) (dao.PostDao, error)
	ArchiveExpiredPosts(ctx context.Context) error
	GetPostsByUserID(userId string, ctx context.Context) ([]dao.PostDao, error)
}
