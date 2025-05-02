package utils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	structures "triple-s/internal/structures"
)

var (
	StorageBase string
	initOnce    sync.Once
)

func CreateCsv(filePath string, headers []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.Write(headers); err != nil {
		return err
	}
	writer.Flush()
	return writer.Error()
}

// InitStorage initializes the base storage directory (thread-safe)
func InitStorage(basePath string) error {
	var initErr error
	initOnce.Do(func() {
		absPath, err := filepath.Abs(basePath)
		if err != nil {
			initErr = err
			return
		}

		if err := os.MkdirAll(absPath, 0o750); err != nil {
			initErr = err
			return
		}

		StorageBase = absPath
	})
	return initErr
}

// StorageDir is the global variable to hold the directory for storing data.
var StorageDir = ""

// SyncDir updates the global StorageDir variable with a new directory path.
func SyncDir(newStorageDir string) {
	StorageDir = newStorageDir
}

// ValidateDirName checks if the provided directory name is valid based on length and pattern.
func ValidateDirName(dirName string) bool {
	// Check length constraints
	if len(dirName) > 63 || len(dirName) < 3 {
		return false
	}
	// Validate against IPv4 address pattern
	matched, _ := regexp.MatchString(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`, dirName)
	if matched {
		return false
	}
	// Check for consecutive dots
	matched, _ = regexp.MatchString(`[\.]{2}`, dirName)
	if matched {
		return false
	}
	// Check for valid characters in directory name
	matched, _ = regexp.MatchString(`^[a-z0-9]+[a-z0-9.-]+[a-z0-9.]+$`, dirName)
	return matched
}

// ValidateObjectName checks if the provided object name is valid based on allowed characters.
func ValidateObjectName(fileName string) bool {
	matched, _ := regexp.MatchString(`^[0-9A-Za-z\!\-\.\*\_\(\)]+$`, fileName)
	return matched
}

// SaveToCsv saves the metadata of a bucket to a CSV file. Creates the file if it doesn't exist.
func SaveToCsv(bucket structures.Bucket) {
	if _, err := os.Stat(StorageDir + "/buckets.csv"); err != nil {
		CreateCsv(StorageDir+"/buckets.csv", []string{"Name", "CreatedTime", "LastModifiedTime", "Status"})
	}
	writer, file, err := createCSVWriter(StorageDir + "/buckets.csv")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	// Prepare data to write
	data := []string{bucket.Name, bucket.CreatedTime, bucket.LastModifiedTime, bucket.Status}
	writeCSVRecord(writer, data)

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV writer:", err)
	}
	defer file.Close()
}

// createCSVWriter opens a file and prepares a CSV writer for it.
func createCSVWriter(fileName string) (*csv.Writer, *os.File, error) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		fmt.Println("Something went wrong")
		return nil, nil, err
	}
	writer := csv.NewWriter(f)
	return writer, f, nil
}

// writeCSVRecord writes a single record to the CSV using the provided writer.
func writeCSVRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Println("Error writing record to CSV:", err)
	}
}

// GetMetadata retrieves all bucket metadata from the CSV file.
func GetMetadata() []structures.Bucket {
	data, err := readCSVFile(StorageDir + "/buckets.csv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	reader, err := parseCSV(data)
	if err != nil {
		fmt.Println("Error creating CSV reader:", err)
		return nil
	}
	buckets := []structures.Bucket{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV data:", err)
			break
		}
		// Append new bucket from record
		buckets = append(buckets, structures.Bucket{Name: record[0], CreatedTime: record[1], LastModifiedTime: record[2], Status: record[3]})
	}
	return buckets
}

// readCSVFile reads the entire content of the specified CSV file.
func readCSVFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// parseCSV creates a CSV reader from the given byte data.
func parseCSV(data []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader, nil
}

// IsExists checks if a bucket with the given name exists and is active.
func IsExists(dirName string) bool {
	buckets := GetMetadata()

	for i := 0; i < len(buckets); i++ {
		if buckets[i].Name == dirName && buckets[i].Status == "Active" {
			return true
		}
	}
	return false
}

// IsEmpty checks if the specified directory is empty.
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(StorageDir + "/" + name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Check if directory has any files
	if err == io.EOF {
		return true, nil
	}
	return false, err // Return false if not empty or if there's an error
}

// RemoveFromMetadata removes a bucket from the metadata by name.
func RemoveFromMetadata(dirName string) {
	buckets := GetMetadata()

	for i := 1; i < len(buckets); i++ {
		if buckets[i].Name == dirName {
			// Remove the bucket and retain the rest
			temp := buckets[i+1:]
			buckets = buckets[:i]
			buckets = append(buckets, temp...)
			break
		}
	}
	// Rewrite the CSV without the removed bucket
	os.Remove(StorageDir + "/buckets.csv")
	for i := 1; i < len(buckets); i++ {
		SaveToCsv(buckets[i])
	}
}

// SaveObjectCsv saves an object's metadata to the specified CSV file.
func SaveObjectCsv(object structures.Object, dirName string) {
	destination := StorageDir + "/" + dirName + "/" + "objects.csv"

	if _, err := os.Stat(destination); err != nil {
		CreateCsv(destination, []string{"ObjectKey", "Size", "ContentType", "LastModified"})
	}
	if IsFileExist(object.ObjectKey, dirName) {
		RewriteFile(object, dirName)
		return
	}

	writer, file, err := createCSVWriter(destination)
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}
	data := []string{object.ObjectKey, object.Size, object.ContentType, object.LastModified}
	writeCSVRecord(writer, data)

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV writer:", err)
	}
	defer file.Close()
}

// GetObjectMetadata retrieves object metadata from the specified CSV file.
func GetObjectMetadata(destination string) []structures.Object {
	data, err := readCSVFile(destination)
	reader, err := parseCSV(data)
	if err != nil {
		fmt.Println("Error creating CSV reader:", err)
		return nil
	}
	objects := []structures.Object{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV data:", err)
			break
		}
		// Append new object from record
		objects = append(objects, structures.Object{ObjectKey: record[0], Size: record[1], ContentType: record[2], LastModified: record[3]})
	}
	return objects
}

// IsFileExist checks if a specific object key exists in the object's CSV file.
func IsFileExist(objectKey string, dirName string) bool {
	objects := GetObjectMetadata(StorageDir + "/" + dirName + "/" + "objects.csv")

	for i := 1; i < len(objects); i++ {
		if objects[i].ObjectKey == objectKey {
			return true
		}
	}
	return false
}

// RewriteFile updates an existing object in the object's CSV file.
func RewriteFile(obj structures.Object, dirName string) {
	objects := GetObjectMetadata(StorageDir + "/" + dirName + "/" + "objects.csv")
	newObjects := []structures.Object{}

	for i := 1; i < len(objects); i++ {
		if objects[i].ObjectKey == obj.ObjectKey {
			// Update object with new data
			newObjects = append(newObjects, obj)
			continue
		}
		newObjects = append(newObjects, objects[i])
	}
	os.Remove(StorageDir + "/" + dirName + "/" + "objects.csv")
	for i := 0; i < len(newObjects); i++ {
		SaveObjectCsv(newObjects[i], dirName)
	}
}

// ObjectMetadata retrieves the metadata of a specific object by its name from the CSV.
func ObjectMetadata(dirName string, fileName string) structures.Object {
	objects := GetObjectMetadata(StorageDir + "/" + dirName + "/" + "objects.csv")

	for i := 1; i < len(objects); i++ {
		if objects[i].ObjectKey == fileName {
			return objects[i] // Return the object if found
		}
	}
	return structures.Object{} // Return an empty object if not found
}

// RemoveObjectFromMetadata removes an object from the metadata by its key.
func RemoveObjectFromMetadata(dirName string, fileName string) {
	objects := GetObjectMetadata(StorageDir + "/" + dirName + "/" + "objects.csv")
	for i := 1; i < len(objects); i++ {
		if objects[i].ObjectKey == fileName {
			// Remove the object and retain the rest
			temp := objects[i+1:]
			objects = objects[:i]
			objects = append(objects, temp...)
			break
		}
	}
	// Rewrite the CSV without the removed object
	os.Remove(StorageDir + "/" + dirName + "/" + "objects.csv")
	for i := 1; i < len(objects); i++ {
		SaveObjectCsv(objects[i], dirName)
	}
}

// UpdateBucket updates the metadata of an existing bucket in the CSV file.
func UpdateBucket(bucket structures.Bucket) {
	buckets := GetMetadata()

	for i := 1; i < len(buckets); i++ {
		if buckets[i].Name == bucket.Name {
			// Preserve existing created time and status while updating
			bucket.CreatedTime = buckets[i].CreatedTime
			bucket.Status = buckets[i].Status
			buckets[i] = bucket // Update the bucket in the slice
		}
	}
	// Rewrite the CSV with updated buckets
	os.Remove(StorageDir + "/buckets.csv")
	for i := 1; i < len(buckets); i++ {
		SaveToCsv(buckets[i])
	}
}
