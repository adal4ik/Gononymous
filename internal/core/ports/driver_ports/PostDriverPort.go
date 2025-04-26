package driverports

import "Gononymous/internal/core/domains/dto"

type PostDriverPortInterface interface {
	AddPost(post dto.PostDto) error
}
