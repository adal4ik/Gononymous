package db

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// Setup a test database connection before running the tests
func setuppp() (*sql.DB, *UserRepo, error) {
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

	userRepo := NewUserRepository(db)

	return db, userRepo, nil
}

// Teardown the database connection after tests
func teardownnn(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// Simple Test: Test ChangeName
func TestChangeName(t *testing.T) {
	db, userRepo, err := setuppp()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardownnn(db)

	// Prepare initial data
	userId := "123"
	initialName := "Old Name"
	newName := "New Name"
	ctx := context.Background()

	// Insert a test user into the database for testing
	_, err = db.ExecContext(ctx, "INSERT INTO users (user_id, name) VALUES ($1, $2)", userId, initialName)
	if err != nil {
		t.Fatalf("Failed to insert test user: %v", err)
	}

	// Call ChangeName to update the user's name
	err = userRepo.ChangeName(userId, newName, ctx)
	if err != nil {
		t.Errorf("Failed to change name: %v", err)
	}

	// Verify that the name was updated
	var updatedName string
	err = db.QueryRowContext(ctx, "SELECT name FROM users WHERE user_id = $1", userId).Scan(&updatedName)
	if err != nil {
		t.Fatalf("Failed to fetch updated name: %v", err)
	}

	// Check if the name was updated correctly
	if updatedName != newName {
		t.Errorf("Expected name to be '%s', got '%s'", newName, updatedName)
	}
}
