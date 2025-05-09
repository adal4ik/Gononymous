package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIResponse_Send(t *testing.T) {
	// Arrange: Create a test APIResponse
	resp := &APIResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}

	// Create a ResponseRecorder to capture the response
	rec := httptest.NewRecorder()

	// Act: Send the response using the Send method
	resp.Send(rec)

	// Assert: Check that the status code is correct
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	// Assert: Check that the content type is set to JSON
	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected content type application/json, got %s", rec.Header().Get("Content-Type"))
	}

	// Assert: Check if the response body contains the expected message
	expected := `{
		"code": 200,
		"message": "Success"
	}`

	if rec.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rec.Body.String())
	}
}

func TestAPIError_Send(t *testing.T) {
	// Arrange: Create a test APIError
	errResp := &APIError{
		Code:     http.StatusBadRequest,
		Message:  "Bad Request",
		Resource: "Post",
	}

	// Create a ResponseRecorder to capture the response
	rec := httptest.NewRecorder()

	// Act: Send the error response using the Send method
	errResp.Send(rec)

	// Assert: Check that the status code is correct
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
	}

	// Assert: Check that the content type is set to JSON
	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected content type application/json, got %s", rec.Header().Get("Content-Type"))
	}

	// Assert: Check if the response body contains the expected message
	expected := `{
		"code": 400,
		"message": "Bad Request",
		"resource": "Post"
	}`

	if rec.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rec.Body.String())
	}
}
