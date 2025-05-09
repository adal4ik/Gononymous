package dao

import (
	"backend/internal/core/domains/dto"
	"time"
)

type PostDao struct {
	PostId     string
	UserId     string
	UserName   string
	UserAvaUrl string
	CreatedAt  time.Time
	Title      string
	Subject    string
	Content    string
	ImageUrl   string
	Status     string
}

func ParseDTOtoDAO(post dto.PostDto) PostDao {
	var newPost PostDao
	newPost.Title = post.Title
	newPost.Content = post.Content
	// newPost.ImageUrl
	newPost.Subject = post.Subject
	return newPost
}
