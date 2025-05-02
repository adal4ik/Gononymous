package main

import (
	"flag"
	"log"
	"os"

	"triple-s/internal"
	"triple-s/internal/utils"
)

func main() {
	// Get storage path (environment variable takes precedence)
	storagePath := getStoragePath()
	port := getPort()

	// Initialize storage
	if err := utils.InitStorage(storagePath); err != nil {
		log.Fatalf("Storage initialization failed: %v", err)
	}

	// Start server
	if err := internal.StartServer(storagePath, port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func getStoragePath() string {
	// Check environment variable first
	if envPath := os.Getenv("STORAGE_PATH"); envPath != "" {
		return envPath
	}

	// Fallback to flag value
	storageFlag := flag.String("dir", "./s3-data", "Storage directory path")
	flag.Parse()
	return *storageFlag
}

func getPort() string {
	// Check environment variable first
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ensureColonPrefix(envPort)
	}

	// Fallback to flag value
	portFlag := flag.String("port", ":9000", "Server port")
	flag.Parse()
	return ensureColonPrefix(*portFlag)
}

func ensureColonPrefix(port string) string {
	if port != "" && port[0] != ':' {
		return ":" + port
	}
	return port
}
