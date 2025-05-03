package test

import (
	"os"
	"testing"

	"backend/internal/adapters/driven/s3"
)

func TestS3SavingImages(t *testing.T) {
	data, err := os.ReadFile("ERD.png")
	if err != nil {
		t.Fatal("Failed to read test image:", err)
	}
	newS3 := s3.NewS3ImageCollector()
	if err != nil {
		t.Fatal("SaveImage failed:", err)
	}
	filePath, err := newS3.SaveImage(data)
	if err != nil {
		t.Fatal("SaveImage failed:", err)
	}
	t.Logf("File successfully created at: %s", filePath)
}
