package driverports

import (
	"Gononymous/internal/core/domains/dto"
	"context"
)

type PostDriverPortInterface interface {
	AddPost(ctx context.Context, post dto.PostDto) error
}
