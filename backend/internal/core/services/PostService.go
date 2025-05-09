package services

import (
	"context"
	"log"
	"time"

	"backend/internal/core/domains/dao"
	"backend/internal/core/domains/dto"
	drivenports "backend/internal/core/ports/driven_ports"
	"backend/utils"
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
	postDao.UserId = post.AuthorID
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

func (postService *PostService) GetPostById(id string) (dto.PostDto, error) {
	postDao, err := postService.repo.GetPostById(id)
	if err != nil {
		return dto.PostDto{}, err
	}
	var postDto dto.PostDto
	postDto.ID = postDao.PostId
	postDto.Image = postDao.ImageUrl
	postDto.Content = postDao.Content
	postDto.Subject = postDao.Subject
	postDto.Title = postDao.Title
	return postDto, nil
}

func (s *PostService) StartPostArchiver(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := s.repo.ArchiveExpiredPosts(ctx)
				if err != nil {
					log.Println("Archiver error:", err)
				}
			}
		}
	}()
}
