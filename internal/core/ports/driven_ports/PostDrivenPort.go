package drivenports

import "Gononymous/internal/core/domains/dao"

type PostDrivenPortInterface interface {
	AddPost(post dao.PostDao) error
}
