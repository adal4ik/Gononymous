package driverports

import "backend/internal/core/domains/dto"

type PostDriverPortInterface interface {
	AddPost(post dto.PostDto, data []byte) error
	GetAll() ([]dto.PostDto, error)
}
