package s3

import (
	"backend/utils"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockS3Server struct{}

func (m *MockS3Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Mock PUT request handling for image upload
	if r.Method == http.MethodPut {
		// Simulate a successful image upload response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("File uploaded successfully"))
		return
	}
	// Handle other request methods if needed
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func TestSaveImage_Success(t *testing.T) {
	// Create a mock HTTP server for the S3 service
	mockServer := httptest.NewServer(&MockS3Server{})
	defer mockServer.Close()

	// Replace the URL in the S3ImageCollector with the mock server's URL
	s3Collector := &S3ImageCollector{
		client: mockServer.Client(),
	}

	// Mock image byte data
	imgData := []byte{0xFF, 0xD8, 0xFF} // A sample JPEG image header

	// Save the image
	uploadedURL, err := s3Collector.SaveImage(imgData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the result
	expectedURL := fmt.Sprintf("%s/images/%s.jpg", mockServer.URL, utils.UUID())
	if !bytes.Contains([]byte(uploadedURL), []byte(expectedURL)) {
		t.Errorf("Expected uploaded URL to contain %v, got %v", expectedURL, uploadedURL)
	}
}

func TestSaveImage_EmptyImage(t *testing.T) {
	// Create a mock HTTP server for the S3 service
	mockServer := httptest.NewServer(&MockS3Server{})
	defer mockServer.Close()

	// Replace the URL in the S3ImageCollector with the mock server's URL
	s3Collector := &S3ImageCollector{
		client: mockServer.Client(),
	}

	// Test with an empty image
	imgData := []byte{}

	// Save the image
	uploadedURL, err := s3Collector.SaveImage(imgData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Ensure that no URL is returned for an empty image
	if uploadedURL != "" {
		t.Errorf("Expected an empty URL for empty image, got %v", uploadedURL)
	}
}

func TestSaveImage_Failure(t *testing.T) {
	// Create a mock HTTP server that simulates failure
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a failure with a bad request
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	// Replace the URL in the S3ImageCollector with the mock server's URL
	s3Collector := &S3ImageCollector{
		client: mockServer.Client(),
	}

	// Mock image byte data
	imgData := []byte{0xFF, 0xD8, 0xFF} // A sample JPEG image header

	// Save the image
	uploadedURL, err := s3Collector.SaveImage(imgData)
	if err == nil {
		t.Fatal("Expected an error, but got none")
	}

	// Check if error is returned
	if uploadedURL != "" {
		t.Errorf("Expected empty URL due to error, got %v", uploadedURL)
	}
}

func TestDetectImageExtension(t *testing.T) {
	// Test JPEG detection
	imgData := []byte{0xFF, 0xD8, 0xFF}
	ext, err := detectImageExtension(imgData)
	if err != nil || ext != ".jpg" {
		t.Errorf("Expected .jpg extension, got %v, error: %v", ext, err)
	}

	// Test PNG detection
	imgData = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	ext, err = detectImageExtension(imgData)
	if err != nil || ext != ".png" {
		t.Errorf("Expected .png extension, got %v, error: %v", ext, err)
	}

	// Test unknown image format
	imgData = []byte{0x00, 0x00, 0x00}
	ext, err = detectImageExtension(imgData)
	if err == nil {
		t.Errorf("Expected error for unknown image format, got extension %v", ext)
	}
}
