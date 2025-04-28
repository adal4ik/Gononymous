package fileHandlers

import (
	"Gononymous/internal/adapters/driven/s3/errs"
	"Gononymous/internal/adapters/driven/s3/structures"
	"Gononymous/internal/adapters/driven/s3/utils"
	"net/http"
	"os"
	"time"
)

// Method for "DELETE" request for objects
func DeleteFileHandler(w http.ResponseWriter, req *http.Request) {
	// Getting bucket name and object key data
	dirName := req.PathValue("BucketName")
	fileName := req.PathValue("ObjectKey")

	// Checking for existense of the bucket
	if !utils.IsExists(dirName) {
		err := errs.Error{Code: 404, Message: "The bucket doesn't exists: " + dirName, Resource: req.URL.Path}
		err.WriteError(w)
		return
	}

	// Checking for existense of the file
	if !utils.IsFileExist(fileName, dirName) {
		err := errs.Error{Code: 404, Message: "The file doesn't exists: " + fileName, Resource: req.URL.Path}
		err.WriteError(w)
		return
	}
	// Creting a new instanse with the bucket information in order to update the corresponding bucket data
	updatedBucket := structures.Bucket{Name: dirName, LastModifiedTime: time.Now().Format(time.RFC1123)}
	utils.UpdateBucket(updatedBucket)
	// Deleting the file
	os.Remove(StorageName + "/" + dirName + "/" + fileName)
	// Removing objects metadata
	utils.RemoveObjectFromMetadata(dirName, fileName)
	w.WriteHeader(204)
}
