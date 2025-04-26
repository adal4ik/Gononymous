package services

import (
	"crypto/rand"
	"fmt"
	"time"

	"Gononymous/internal/core/domains/dao"
	"Gononymous/internal/core/domains/dto"
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
	postDao.PostId = pseudo_uuid()
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
