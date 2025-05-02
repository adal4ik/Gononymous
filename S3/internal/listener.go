package internal

import (
	"fmt"
	"log"
	"net/http"

	buckets "triple-s/internal/bucketHandlers"
	"triple-s/internal/fileHandlers"
	"triple-s/internal/utils"
)

func StartServer(storagePath, port string) error {
	// 1. Initialize storage directory
	if err := utils.InitStorage(storagePath); err != nil {
		return fmt.Errorf("failed to initialize storage: %w", err)
	}

	// 2. Initialize handlers
	if err := buckets.Init(); err != nil {
		return fmt.Errorf("bucket handler initialization failed: %w", err)
	}
	if err := fileHandlers.Init(); err != nil {
		return fmt.Errorf("file handler initialization failed: %w", err)
	}

	// 3. Configure routes
	router := http.NewServeMux()

	// Bucket operations
	router.HandleFunc("PUT /{BucketName}", buckets.PutHandler)
	router.HandleFunc("GET /", buckets.GetHandler)
	router.HandleFunc("DELETE /{BucketName}", buckets.DeleteHandler)

	// Object operations
	router.HandleFunc("PUT /{BucketName}/{ObjectKey}", fileHandlers.PutFileHandler)
	router.HandleFunc("GET /{BucketName}/{ObjectKey}", fileHandlers.GetFileHandler)
	router.HandleFunc("DELETE /{BucketName}/{ObjectKey}", fileHandlers.DeleteFileHandler)

	// Health check
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// 4. Start server
	log.Printf("Server starting on port %s", port)
	log.Printf("Storage directory: %s", storagePath)

	if err := http.ListenAndServe(port, router); err != nil {
		return fmt.Errorf("server failed: %w", err)
	}

	return nil
}
