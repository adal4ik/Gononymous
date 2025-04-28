package fileHandlers

import (
	"Gononymous/internal/adapters/driven/s3/errs"
	"Gononymous/internal/adapters/driven/s3/structures"
	"Gononymous/internal/adapters/driven/s3/utils"
	"io"
	"net/http"
	"os"
	"time"
)

var StorageName = "Gononymous/internal/adapters/driven/s3"

// Method for "PUT" request for objects
func PutFileHandler(w http.ResponseWriter, req *http.Request) {
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
	if !utils.ValidateObjectName(fileName) {
		err := errs.Error{Code: 404, Message: "The file doesn't exists: " + fileName, Resource: req.URL.Path}
		err.WriteError(w)
		return
	}
	// Creting a new instanse with the bucket information in order to update the corresponding bucket data
	updatedBucket := structures.Bucket{Name: dirName, LastModifiedTime: time.Now().Format(time.RFC1123)}
	utils.UpdateBucket(updatedBucket)
	// Creating the file
	out, _ := os.Create(StorageName + "/" + dirName + "/" + fileName)
	// Copying the data from the request body to the file
	io.Copy(out, req.Body)
	// Saving objects's metadata
	object := structures.Object{ObjectKey: fileName, Size: req.Header.Get("Content-Length"), ContentType: req.Header.Get("Content-Type"), LastModified: time.Now().Format(time.RFC1123)}
	utils.SaveObjectCsv(object, dirName)
}
