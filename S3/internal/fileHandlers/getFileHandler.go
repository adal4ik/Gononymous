package fileHandlers

import (
	"io"
	"net/http"
	"os"

	"triple-s/internal/utils"
)

// Method for "GET" request for objects
func GetFileHandler(w http.ResponseWriter, req *http.Request) {
	// Getting bucket name and object key data
	dirName := req.PathValue("BucketName")
	fileName := req.PathValue("ObjectKey")

	// Checking for existense of the bucket
	// if !utils.IsExists(dirName) {
	// 	err := errs.Error{Code: 404, Message: "The bucket doesn't exists: " + dirName, Resource: req.URL.Path}
	// 	err.WriteError(w)
	// 	return
	// }
	// Checking for existense of the file
	// if !utils.IsFileExist(fileName, dirName) {
	// 	err := errs.Error{Code: 404, Message: "The file doesn't exists: " + fileName, Resource: req.URL.Path}
	// 	err.WriteError(w)
	// 	return
	// }
	// Creting a new instanse with the bucket information in order to update the corresponding bucket data
	// updatedBucket := structures.Bucket{Name: dirName, LastModifiedTime: time.Now().Format(time.RFC1123)}
	// utils.UpdateBucket(updatedBucket)
	// Getting the object metadata
	theObject := utils.ObjectMetadata(dirName, fileName)
	// Openning the file
	file, _ := os.Open(StorageName + "/" + dirName + "/" + fileName)
	// Writing the metadata to the header
	w.Header().Set("Content-Length", theObject.Size)
	w.Header().Set("Content-Type", theObject.ContentType)
	// Retriving the file
	io.Copy(w, file)
}
