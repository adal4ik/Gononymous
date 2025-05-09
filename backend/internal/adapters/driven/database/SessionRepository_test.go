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
func setupp() (*sql.DB, *SessionRepo, error) {
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

	sessionRepo := NewSessionRepo(db)

	return db, sessionRepo, nil
}

// Teardown the database connection after tests
func teardownn(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// Simple Test: Test AddSession and GetSessionById
func TestSessionRepo(t *testing.T) {
	db, sessionRepo, err := setupp()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardownn(db)

	// Prepare test data
	session := dao.Session{
		UsersId:   "1",
		Name:      "Test User",
		AvatarURL: "http://example.com/avatar.jpg",
	}

	ctx := context.Background()

	// Test AddSession
	err = sessionRepo.AddSession(ctx, session)
	if err != nil {
		t.Errorf("Failed to add session: %v", err)
	}

	// Test GetSessionById
	retrievedSession, err := sessionRepo.GetSessionById(session.UsersId, ctx)
	if err != nil {
		t.Errorf("Failed to retrieve session: %v", err)
	}

	// Check if the session retrieved is the same as the added session
	if retrievedSession.UsersId != session.UsersId {
		t.Errorf("Session ID mismatch: expected %s, got %s", session.UsersId, retrievedSession.UsersId)
	}
	if retrievedSession.Name != session.Name {
		t.Errorf("Session Name mismatch: expected %s, got %s", session.Name, retrievedSession.Name)
	}
	if retrievedSession.AvatarURL != session.AvatarURL {
		t.Errorf("Session AvatarURL mismatch: expected %s, got %s", session.AvatarURL, retrievedSession.AvatarURL)
	}
}
