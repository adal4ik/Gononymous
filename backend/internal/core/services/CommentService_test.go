package services

import (
	"backend/internal/core/domains/dto"
	"context"
	"fmt"
	"testing"
)

// Mock for CommentRepoInterface
type MockCommentRepo struct{}

func (m *MockCommentRepo) AddComment(comment dto.Comment, ctx context.Context) error {
	// Simulate adding a comment
	if comment.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}
	return nil
}

func (m *MockCommentRepo) GetCommentsByPostId(postId string, ctx context.Context) ([]dto.Comment, error) {
	// Simulate getting comments by post ID
	return []dto.Comment{
		{CommentID: "1", PostID: postId, Content: "Comment 1"},
		{CommentID: "2", PostID: postId, Content: "Comment 2"},
	}, nil
}

func (m *MockCommentRepo) GetCommentReplies(commentId string, ctx context.Context) ([]dto.Comment, error) {
	// Simulate getting replies for a comment
	if commentId == "1" {
		return []dto.Comment{
			{CommentID: "3", PostID: "1", ParentID: "1", Content: "Reply to Comment 1"},
		}, nil
	}
	return nil, nil
}

// Mock for ImageCollectionInterface
type MockImageCollector struct{}

func (m *MockImageCollector) SaveImage(img []byte) (string, error) {
	// Simulate saving an image
	if len(img) == 0 {
		return "", fmt.Errorf("image is empty")
	}
	return "http://example.com/image.jpg", nil
}

func TestAddComment(t *testing.T) {
	// Arrange
	mockRepo := &MockCommentRepo{}
	mockImageCollector := &MockImageCollector{}

	commentService := NewCommentService(mockRepo, mockImageCollector)

	comment := dto.Comment{
		PostID:  "1",
		UserID:  "user1",
		Content: "This is a comment",
	}
	img := []byte{0xFF, 0xD8} // Some mock image data

	// Act
	err := commentService.AddComment(comment, img, context.Background())
	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAddComment_ImageError(t *testing.T) {
	// Arrange
	mockRepo := &MockCommentRepo{}
	mockImageCollector := &MockImageCollector{}

	commentService := NewCommentService(mockRepo, mockImageCollector)

	comment := dto.Comment{
		PostID:  "1",
		UserID:  "user1",
		Content: "This is a comment",
	}
	img := []byte{} // Simulating an empty image

	// Act
	err := commentService.AddComment(comment, img, context.Background())

	// Assert
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGetCommentsByPostId(t *testing.T) {
	// Arrange
	mockRepo := &MockCommentRepo{}
	mockImageCollector := &MockImageCollector{}

	commentService := NewCommentService(mockRepo, mockImageCollector)

	// Act
	comments, err := commentService.GetCommentsByPostId("1", context.Background())
	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(comments) != 2 {
		t.Errorf("Expected 2 comments, got %v", len(comments))
	}

	if len(comments[0].Replies) != 1 {
		t.Errorf("Expected 1 reply for comment 1, got %v", len(comments[0].Replies))
	}

	if len(comments[1].Replies) != 0 {
		t.Errorf("Expected 0 replies for comment 2, got %v", len(comments[1].Replies))
	}
}
