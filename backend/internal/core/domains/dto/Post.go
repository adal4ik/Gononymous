package dto

import "time"

type PostDto struct {
	ID           string
	AuthorID     string
	AuthorName   string
	AuthorAvaUrl string
	Title        string
	Subject      string
	Content      string
	Image        string
	CreatedAt    time.Time
}
