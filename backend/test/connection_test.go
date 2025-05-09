package test

import (
	db "backend/internal/adapters/driven/database"
	"log"
	"os"
	"testing"
)

// Mock ConnectDB using real database interaction
func TestConnectDB(t *testing.T) {
	// Set environment variables for the test (or use real test values)
	os.Setenv("DB_HOST", "localhost")        // Your test DB host
	os.Setenv("DB_USER", "testuser")         // Your test DB user
	os.Setenv("DB_PASSWORD", "testpassword") // Your test DB password
	os.Setenv("DB_NAME", "testdb")           // Your test DB name
	os.Setenv("DB_PORT", "5432")             // Your test DB port

	// Call ConnectDB to test
	db := db.ConnectDB()

	// Check if the connection is established
	if db == nil {
		t.Fatalf("Failed to connect to the database!")
	}

	// If we reach here, the connection and ping succeeded
	log.Println("Successfully connected to the test database!")
}
