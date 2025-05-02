package services

import (
	"time"

	"backend/internal/core/domains/dao"
	"backend/internal/core/domains/dto"
	"backend/utils"

	drivenports "backend/internal/core/ports/driven_ports"
)

type PostService struct {
	repo           drivenports.DatabasePortInterface
	imageCollector drivenports.ImageCollectionInterface
}

func NewPostService(repo drivenports.DatabasePortInterface, imageCollector drivenports.ImageCollectionInterface) *PostService {
	return &PostService{repo: repo, imageCollector: imageCollector}
}

func (postService *PostService) AddPost(post dto.PostDto, data []byte) error {
	var err error
	postDao := dao.ParseDTOtoDAO(post)
	postDao.CreatedAt = time.Now()
	postDao.PostId = utils.UUID()
	postDao.Status = "Active"
	postDao.ImageUrl, err = postService.imageCollector.SaveImage(data)
	if err != nil {
		return err
	}
	err = postService.repo.AddPost(postDao)
	if err != nil {
		return err
	}
	return nil
}

func (postService *PostService) GetAll() ([]dto.PostDto, error) {
	posts, err := postService.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var postsDTO []dto.PostDto

	for i := 0; i < len(posts); i++ {
		postsDTO = append(postsDTO, dto.PostDto{ID: posts[i].PostId, Title: posts[i].Title, Subject: posts[i].Subject, Content: posts[i].Content, Image: posts[i].ImageUrl})
	}
	return postsDTO, nil
}
