package drivenports

import (
	"Gononymous/internal/core/domains/dao"
	"context"
)

type PostDrivenPortInterface interface {
	AddPost(ctx context.Context, post dao.PostDao) error
}
