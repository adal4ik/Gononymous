package dao

import (
	"time"

	"Gononymous/internal/core/domains/dto"
)

type PostDao struct {
	PostId    string
	UserId    string
	CreatedAt time.Time
	Title     string
	Subject   string
	Content   string
	ImageUrl  string
}

func ParseDTOtoDAO(post dto.PostDto) PostDao {
	var newPost PostDao
	newPost.Title = post.Title
	newPost.Content = post.Content
	// newPost.ImageUrl
	newPost.Subject = post.Subject
	return newPost
}
