package services

import (
	"backend/internal/core/domains/dto"
	drivenports "backend/internal/core/ports/driven_ports"
	"backend/utils"
	"fmt"
)

type CommentService struct {
	repo           drivenports.CommentRepoInterface
	imageCollector drivenports.ImageCollectionInterface
}

func NewCommentService(repo drivenports.CommentRepoInterface, imageCollector drivenports.ImageCollectionInterface) *CommentService {
	return &CommentService{repo: repo, imageCollector: imageCollector}
}

func (commentService *CommentService) AddComment(comment dto.Comment, img []byte) error {
	var err error
	comment.CommentID = utils.UUID()
	if len(comment.ParentID) == 0 {
		comment.ParentID = "00000000-0000-0000-0000-000000000000"
	}
	comment.ImageUrl, err = commentService.imageCollector.SaveImage(img)
	if err != nil {
		return err
	}
	err = commentService.repo.AddComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (commentService *CommentService) GetCommentsByPostId(postId string) ([]dto.Comment, error) {
	comments, err := commentService.repo.GetCommentsByPostId(postId)
	for i := 0; i < len(comments); i++ {
		comments[i].Replies, err = commentService.repo.GetCommentReplies(comments[i].CommentID)
		fmt.Println(comments[i].Replies)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return comments, nil
}
