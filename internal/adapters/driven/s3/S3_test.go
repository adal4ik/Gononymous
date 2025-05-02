package s3

import (
	"os"
	"testing"
)

func TestS3SavingImages(t *testing.T) {
	data, err := os.ReadFile("MainBefore.jpg")
	if err != nil {
		t.Fatal("Failed to read test image:", err)
	}
	newS3 := NewS3ImageCollector()
	filePath, err := newS3.SaveImage(data)
	if err != nil {
		t.Fatal("SaveImage failed:", err)
	}

	// Verify the file was created
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("File was not created at: %s", filePath)
	} else {
		t.Logf("File successfully created at: %s", filePath)
	}
}
