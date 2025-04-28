package drivenports

import "Gononymous/internal/core/domains/dao"

type DatabasePortInterface interface {
	AddPost(post dao.PostDao) error
	GetAll() ([]dao.PostDao, error)
}
