package externalapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCharacter_Success(t *testing.T) {
	// Create a mock server to simulate a successful API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ensure the correct URL is requested
		if r.URL.Path != "/api/character/1" {
			t.Errorf("Expected request to '/api/character/1', got '%s'", r.URL.Path)
		}

		// Respond with a mock character data
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"id": 1,
			"name": "Rick Sanchez",
			"status": "Alive",
			"species": "Human",
			"type": "Scientist",
			"gender": "Male"
		}`))
	}))
	defer mockServer.Close()

	// Create the client and inject the mock server's URL
	client := &CharacterClient{
		httpClient: mockServer.Client(),
	}

	// Call the method
	ctx := context.Background()
	character, err := client.GetCharacter(ctx, 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate the response
	if character == nil {
		t.Fatalf("Expected character to be not nil")
	}
	if character.Name != "Rick Sanchez" {
		t.Errorf("Expected character name 'Rick Sanchez', got '%s'", character.Name)
	}
}

func TestGetCharacter_Failure_RequestError(t *testing.T) {
	// Create a mock server to simulate a failure in the request
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a failure in the server
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	// Create the client and inject the mock server's URL
	client := &CharacterClient{
		httpClient: mockServer.Client(),
	}

	// Call the method
	ctx := context.Background()
	_, err := client.GetCharacter(ctx, 1)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if err.Error() != "unexpected status: 500 Internal Server Error" {
		t.Errorf("Expected error message 'unexpected status: 500 Internal Server Error', got %v", err)
	}
}

func TestGetCharacter_Failure_JsonDecodeError(t *testing.T) {
	// Create a mock server to simulate an invalid JSON response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Send an invalid JSON response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{ "invalid": "json" `)) // Missing closing brace
	}))
	defer mockServer.Close()

	// Create the client and inject the mock server's URL
	client := &CharacterClient{
		httpClient: mockServer.Client(),
	}

	// Call the method
	ctx := context.Background()
	_, err := client.GetCharacter(ctx, 1)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if err.Error() != "failed to decode character: unexpected EOF" {
		t.Errorf("Expected error message 'failed to decode character: unexpected EOF', got %v", err)
	}
}

func TestGetCharacter_Failure_RequestCreationError(t *testing.T) {
	// Create the client with an invalid configuration that causes a request creation error
	client := &CharacterClient{
		httpClient: &http.Client{}, // Using an empty client to simulate an error
	}

	// Call the method with an invalid context or bad request
	ctx := context.Background()
	_, err := client.GetCharacter(ctx, -1) // Negative ID to simulate bad URL creation
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if err.Error() != "failed to create request: net/http: invalid method \"GET /api/character/-1\" and URL \"\" in request" {
		t.Errorf("Expected request creation error, got %v", err)
	}
}
