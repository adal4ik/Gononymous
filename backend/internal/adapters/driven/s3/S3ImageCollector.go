package s3

import (
	"backend/utils"
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"path/filepath"
)

type S3ImageCollector struct {
	client *http.Client
}

func NewS3ImageCollector() *S3ImageCollector {
	req, _ := http.NewRequest("PUT", "http://s3:9000/images", nil)
	customCLinet := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	customCLinet.Do(req)
	return &S3ImageCollector{client: customCLinet}
}

func (s *S3ImageCollector) SaveImage(img []byte) (string, error) {
	if len(img) == 0 {
		return "", nil
	}
	bucket := "images/"

	ext, err := detectImageExtension(img)
	if err != nil {
		return "", fmt.Errorf("could not detect image type: %w", err)
	}

	fileName := utils.UUID() + ext
	filePath := filepath.Join(bucket, fileName)

	// Create a new request
	req, err := http.NewRequest("PUT", "http://s3:9000/"+filePath, bytes.NewReader(img))
	if err != nil {
		return "", fmt.Errorf("could not create request: %w", err)
	}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/octet-stream")

	// Send the request
	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("server returned error status: %s", resp.Status)
	}

	return "http://localhost:9000/" + bucket + fileName, nil
}

func detectImageExtension(img []byte) (string, error) {
	if len(img) < 8 {
		return "", fmt.Errorf("file too small to determine type")
	}

	if bytes.HasPrefix(img, []byte{0xFF, 0xD8}) {
		return ".jpg", nil
	}

	// PNG: Starts with 0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A
	if bytes.HasPrefix(img, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return ".png", nil
	}

	// GIF: Starts with "GIF87a" or "GIF89a"
	if bytes.HasPrefix(img, []byte("GIF87a")) || bytes.HasPrefix(img, []byte("GIF89a")) {
		return ".gif", nil
	}

	// WebP: Starts with "RIFF" followed by "WEBP"
	if len(img) >= 12 && bytes.HasPrefix(img, []byte("RIFF")) && bytes.HasPrefix(img[8:12], []byte("WEBP")) {
		return ".webp", nil
	}

	// BMP: Starts with "BM"
	if bytes.HasPrefix(img, []byte("BM")) {
		return ".bmp", nil
	}

	return "", fmt.Errorf("unknown image format")
}
