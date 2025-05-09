package drivenports

import (
	"context"

	"backend/internal/core/domains/dao"
)

type DatabasePortInterface interface {
	AddPost(post dao.PostDao) error
	GetAll() ([]dao.PostDao, error)
	GetPostById(id string) (dao.PostDao, error)
	ArchiveExpiredPosts(ctx context.Context) error
}
