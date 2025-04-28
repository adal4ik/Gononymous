package services

import (
	"Gononymous/internal/core/domains/dao"
	"Gononymous/internal/core/domains/dto"
	"crypto/rand"
	"fmt"
	"time"

	drivenports "Gononymous/internal/core/ports/driven_ports"
)

type PostService struct {
	repo drivenports.DatabasePortInterface
}

func NewPostService(repo drivenports.DatabasePortInterface) *PostService {
	return &PostService{repo: repo}
}

func (postService *PostService) AddPost(post dto.PostDto) error {
	postDao := dao.ParseDTOtoDAO(post)
	postDao.CreatedAt = time.Now()
	postDao.PostId = pseudo_uuid()
	postDao.Status = "Active"
	err := postService.repo.AddPost(postDao)
	if err != nil {
		return err
	}
	return nil
}

func pseudo_uuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

func (postService *PostService) GetAll() ([]dto.PostDto, error) {
	allPostsDao, err := postService.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var allPostsDto []dto.PostDto

	for _, val := range allPostsDao {
		allPostsDto = append(allPostsDto, dto.PostDto{Title: val.Title, Subject: val.Subject, Content: val.Content})
	}
	return allPostsDto, nil
}
