package services

import (
	"Gononymous/internal/core/domains/dao"
	"Gononymous/internal/core/domains/dto"
	"Gononymous/utils"
	"context"
	"time"

	drivenports "Gononymous/internal/core/ports/driven_ports"
)

type PostService struct {
	repo drivenports.PostDrivenPortInterface
}

func NewPostService(repo drivenports.PostDrivenPortInterface) *PostService {
	return &PostService{repo: repo}
}

func (postService *PostService) AddPost(ctx context.Context, post dto.PostDto) error {
	postDao := dao.ParseDTOtoDAO(post)
	postDao.CreatedAt = time.Now()
	postDao.PostId = utils.UUID()
	err := postService.repo.AddPost(ctx, postDao)
	if err != nil {
		return err
	}
	return nil
}
