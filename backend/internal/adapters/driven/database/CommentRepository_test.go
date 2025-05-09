package db

import (
	"backend/internal/core/domains/dto"
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// Setup a test database connection before running the tests
func setupppp() (*sql.DB, *CommentRepository, error) {
	// Set the environment variables for testing
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "latte")
	os.Setenv("DB_PASSWORD", "latte")
	os.Setenv("DB_NAME", "frappuccino")
	os.Setenv("DB_PORT", "5432")

	// Create the PostgreSQL connection string
	psqlInfo := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, nil, err
	}

	// Ensure the connection is valid
	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	commentRepo := NewCommentRepository(db)

	return db, commentRepo, nil
}

// Teardown the database connection after tests
func teardownnnn(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// Simple Test: Test AddComment, GetCommentsByPostId, and GetCommentReplies
func TestCommentRepository(t *testing.T) {
	db, commentRepo, err := setupppp()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardownnnn(db)

	// Prepare test data
	comment := dto.Comment{
		CommentID: "comment1",
		PostID:    "post1",
		ParentID:  "00000000-0000-0000-0000-000000000000", // Root comment
		UserID:    "user1",
		Content:   "This is a test comment",
		ImageUrl:  "http://example.com/image.jpg",
	}

	ctx := context.Background()

	// Test AddComment
	err = commentRepo.AddComment(comment, ctx)
	if err != nil {
		t.Errorf("Failed to add comment: %v", err)
	}

	// Test GetCommentsByPostId
	comments, err := commentRepo.GetCommentsByPostId(comment.PostID, ctx)
	if err != nil {
		t.Errorf("Failed to get comments by post ID: %v", err)
	}
	if len(comments) == 0 {
		t.Errorf("Expected to find at least one comment, found none")
	}

	// Verify that the comment is returned

	// Test GetCommentReplies (should return an empty list in this case)
	replies, err := commentRepo.GetCommentReplies(comment.CommentID, ctx)
	if err != nil {
		t.Errorf("Failed to get comment replies: %v", err)
	}
	if len(replies) != 0 {
		t.Errorf("Expected no replies for comment ID '%s', got %d replies", comment.CommentID, len(replies))
	}
}
