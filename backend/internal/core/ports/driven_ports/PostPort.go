package drivenports

import (
	"backend/internal/core/domains/dao"
	"context"
)

type DatabasePortInterface interface {
	AddPost(post dao.PostDao) error
	GetActive() ([]dao.PostDao, error)
	GetAll() ([]dao.PostDao, error)
	GetPostById(id string) (dao.PostDao, error)
	ArchiveExpiredPosts(ctx context.Context) error
	GetPostsByUserID(userId string) ([]dao.PostDao, error)
}
