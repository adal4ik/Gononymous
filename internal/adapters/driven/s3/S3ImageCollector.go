package s3

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"Gononymous/utils"
)

type S3ImageCollector struct{}

func NewS3ImageCollector() *S3ImageCollector {
	return &S3ImageCollector{}
}

func (s *S3ImageCollector) SaveImage(img []byte) (string, error) {
	dirName := "Images"

	// Detect file extension
	ext, err := detectImageExtension(img)
	if err != nil {
		return "", fmt.Errorf("could not detect image type: %w", err)
	}

	fileName := utils.UUID() + ext

	if err := os.MkdirAll(dirName, 0o755); err != nil {
		return "", err
	}

	filePath := filepath.Join(dirName, fileName)
	if err := os.WriteFile(filePath, img, 0o666); err != nil {
		return "", err
	}

	return filePath, nil
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
