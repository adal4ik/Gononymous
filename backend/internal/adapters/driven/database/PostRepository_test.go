package db

import (
	"backend/internal/core/domains/dao"
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// Setup a test database connection before running the tests
func setup() (*sql.DB, *PostRepository, error) {
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

	postRepository := NewPostRepository(db)

	return db, postRepository, nil
}

// Teardown the database connection after tests
func teardown(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// TestAddPost tests the AddPost method of the PostRepository
func TestAddPost(t *testing.T) {
	db, postRepository, err := setup()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardown(db)

	// Prepare test data
	post := dao.PostDao{
		PostId:   "1",
		UserId:   "1",
		Title:    "Test Post",
		Subject:  "Test Subject",
		Content:  "Test Content",
		ImageUrl: "http://example.com/image.jpg",
		Status:   "Active",
	}

	ctx := context.Background()

	// Add the post
	err = postRepository.AddPost(post, ctx)
	if err != nil {
		t.Errorf("Failed to add post: %v", err)
	}

	// Verify if the post is in the database by fetching it
	storedPost, err := postRepository.GetPostById(post.PostId, ctx)
	if err != nil {
		t.Errorf("Failed to retrieve post: %v", err)
	}
	if storedPost.PostId != post.PostId {
		t.Errorf("Post ID mismatch: expected %s, got %s", post.PostId, storedPost.PostId)
	}
}

// TestGetActive tests the GetActive method of the PostRepository
func TestGetActive(t *testing.T) {
	db, postRepository, err := setup()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardown(db)

	ctx := context.Background()

	// Add a post to be fetched
	post := dao.PostDao{
		PostId:   "2",
		UserId:   "1",
		Title:    "Active Post",
		Subject:  "Active Subject",
		Content:  "Active Content",
		ImageUrl: "http://example.com/image.jpg",
		Status:   "Active",
	}
	err = postRepository.AddPost(post, ctx)
	if err != nil {
		t.Errorf("Failed to add post: %v", err)
	}

	// Fetch active posts
	activePosts, err := postRepository.GetActive(ctx)
	if err != nil {
		t.Errorf("Failed to get active posts: %v", err)
	}
	if len(activePosts) == 0 {
		t.Error("Expected to get active posts, but got none")
	}
}

// TestGetAll tests the GetAll method of the PostRepository
func TestGetAll(t *testing.T) {
	db, postRepository, err := setup()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardown(db)

	ctx := context.Background()

	// Add test posts
	post1 := dao.PostDao{
		PostId:   "3",
		UserId:   "1",
		Title:    "Post 1",
		Subject:  "Subject 1",
		Content:  "Content 1",
		ImageUrl: "http://example.com/image1.jpg",
		Status:   "Active",
	}
	post2 := dao.PostDao{
		PostId:   "4",
		UserId:   "1",
		Title:    "Post 2",
		Subject:  "Subject 2",
		Content:  "Content 2",
		ImageUrl: "http://example.com/image2.jpg",
		Status:   "Active",
	}
	err = postRepository.AddPost(post1, ctx)
	if err != nil {
		t.Errorf("Failed to add post: %v", err)
	}
	err = postRepository.AddPost(post2, ctx)
	if err != nil {
		t.Errorf("Failed to add post: %v", err)
	}

	// Fetch all posts
	allPosts, err := postRepository.GetAll(ctx)
	if err != nil {
		t.Errorf("Failed to get all posts: %v", err)
	}
	if len(allPosts) < 2 {
		t.Errorf("Expected at least 2 posts, got %d", len(allPosts))
	}
}

// TestGetPostById tests the GetPostById method of the PostRepository
func TestGetPostById(t *testing.T) {
	db, postRepository, err := setup()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardown(db)

	ctx := context.Background()

	// Add a post
	post := dao.PostDao{
		PostId:   "5",
		UserId:   "1",
		Title:    "Test Get Post By ID",
		Subject:  "Test Subject",
		Content:  "Test Content",
		ImageUrl: "http://example.com/image.jpg",
		Status:   "Active",
	}
	err = postRepository.AddPost(post, ctx)
	if err != nil {
		t.Errorf("Failed to add post: %v", err)
	}

	// Fetch the post by ID
	storedPost, err := postRepository.GetPostById(post.PostId, ctx)
	if err != nil {
		t.Errorf("Failed to retrieve post: %v", err)
	}
	if storedPost.PostId != post.PostId {
		t.Errorf("Post ID mismatch: expected %s, got %s", post.PostId, storedPost.PostId)
	}
}

// TestArchiveExpiredPosts tests the ArchiveExpiredPosts method of the PostRepository
func TestArchiveExpiredPosts(t *testing.T) {
	db, postRepository, err := setup()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardown(db)

	ctx := context.Background()

	// Add expired posts
	post := dao.PostDao{
		PostId:   "6",
		UserId:   "1",
		Title:    "Expired Post",
		Subject:  "Expired Subject",
		Content:  "Expired Content",
		ImageUrl: "http://example.com/expired.jpg",
		Status:   "Active",
	}

	err = postRepository.AddPost(post, ctx)

	// Archive expired posts (for testing, simulate time)
	err = postRepository.ArchiveExpiredPosts(ctx)
	if err != nil {
		t.Errorf("Failed to archive expired posts: %v", err)
	}

	// Fetch the post again and check status
	storedPost, err := postRepository.GetPostById(post.PostId, ctx)
	if err != nil {
		t.Errorf("Failed to retrieve post: %v", err)
	}
	if storedPost.Status != "Archived" {
		t.Errorf("Expected post status to be Archived, got %s", storedPost.Status)
	}
}

// TestGetPostsByUserID tests the GetPostsByUserID method of the PostRepository
func TestGetPostsByUserID(t *testing.T) {
	db, postRepository, err := setup()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardown(db)

	ctx := context.Background()

	// Add posts for a user
	post := dao.PostDao{
		PostId:   "7",
		UserId:   "1",
		Title:    "User Post",
		Subject:  "User Subject",
		Content:  "User Content",
		ImageUrl: "http://example.com/userpost.jpg",
		Status:   "Active",
	}

	err = postRepository.AddPost(post, ctx)
	if err != nil {
		t.Errorf("Failed to add post: %v", err)
	}

	// Fetch posts by user ID
	posts, err := postRepository.GetPostsByUserID("1", ctx)
	if err != nil {
		t.Errorf("Failed to get posts by user ID: %v", err)
	}
	if len(posts) == 0 {
		t.Errorf("Expected at least 1 post for user 1, got %d", len(posts))
	}
}
