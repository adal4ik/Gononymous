package services

import (
	"Gononymous/internal/core/domains/dao"
	"Gononymous/internal/core/domains/dto"
	"Gononymous/utils"
	"time"

	drivenports "Gononymous/internal/core/ports/driven_ports"
)

type PostService struct {
	repo drivenports.PostDrivenPortInterface
}

func NewPostService(repo drivenports.PostDrivenPortInterface) *PostService {
	return &PostService{repo: repo}
}

func (postService *PostService) AddPost(post dto.PostDto) error {
	postDao := dao.ParseDTOtoDAO(post)
	postDao.CreatedAt = time.Now()
	postDao.PostId = utils.UUID()
	err := postService.repo.AddPost(postDao)
	if err != nil {
		return err
	}
	return nil
}
