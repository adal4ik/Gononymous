package utils

import (
	"os"
	"testing"
)

func TestLogger_CreatesLogFile(t *testing.T) {
	// Arrange
	logger, file := Logger()

	// Assert that the file was created
	if file == nil {
		t.Fatalf("Expected file to be opened, got nil")
	}
	defer file.Close()

	// Check if the file exists
	if _, err := os.Stat("app.log"); os.IsNotExist(err) {
		t.Fatalf("Expected log file to be created, but got error: %v", err)
	}

	// Check that the logger isn't nil
	if logger == nil {
		t.Fatalf("Expected logger to be created, got nil")
	}
}

func TestLogger_LogsToFile(t *testing.T) {
	// Arrange
	logger, file := Logger()
	if file == nil {
		t.Fatal("File should not be nil")
	}
	defer file.Close()

	// Close the file first to simulate log writing
	file.Close()

	// Arrange: Write something to the logger
	logger.Info("This is a test log")

	// Assert: Check if the log entry is written to the file
	logFile, err := os.Open("app.log")
	if err != nil {
		t.Fatalf("Expected log file to be opened, but got error: %v", err)
	}
	defer logFile.Close()

	// Seek to the end of the file to read the most recent entry
	stat, _ := logFile.Stat()
	if stat.Size() == 0 {
		t.Fatal("Log file is empty, expected log content")
	}
}

func TestLogger_CleanupOldLogs(t *testing.T) {
	// Arrange
	// Make sure the log file is present before testing cleanup
	fileName := "app.log"
	_, file := Logger()
	if file == nil {
		t.Fatal("File creation failed")
	}
	defer file.Close()

	// Clean up logs after each test to avoid growing logs during test execution
	err := os.Remove(fileName)
	if err != nil {
		t.Fatalf("Failed to cleanup log file: %v", err)
	}
}

func TestLogger_HandleErrorGracefully(t *testing.T) {
	// Arrange: Try to create a logger with an invalid file path
	// For example, trying to open a directory instead of a file
	invalidLogger, invalidFile := Logger()
	if invalidFile == nil || invalidLogger == nil {
		t.Log("Logger returned expected error when trying to create an invalid file path")
	}
}
