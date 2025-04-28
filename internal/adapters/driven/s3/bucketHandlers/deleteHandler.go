package buckets

import (
	"Gononymous/internal/adapters/driven/s3/errs"
	"Gononymous/internal/adapters/driven/s3/utils"
	"net/http"
	"os"
)

// Method for "DELETE" request
func DeleteHandler(w http.ResponseWriter, req *http.Request) {
	// Getting a bucket name
	dirName := req.PathValue("BucketName")

	// Cheking for storage existense
	if _, err := os.Stat(StorageName + "/" + "buckets.csv"); err != nil {
		err := errs.Error{Code: 404, Message: "The storage is empty, put something before deleting", Resource: req.URL.Path}
		err.WriteError(w)
		return
	}
	// Checking for existene
	if !utils.IsExists(dirName) {
		err := errs.Error{Code: 404, Message: "The bucket doesn't exists: " + dirName, Resource: req.URL.Path}
		err.WriteError(w)
		return
	}

	// Cheking for emptyness of bucket
	if status, _ := utils.IsEmpty(dirName); !status {
		err := errs.Error{Code: 409, Message: "Conflict for a non-empty bucket: " + dirName, Resource: req.URL.Path}
		err.WriteError(w)
		return
	}
	// Removing the bucket
	os.Remove(StorageName + "/" + dirName)
	// Removing the metadata of the bucket
	utils.RemoveFromMetadata(dirName)
	// Retriving the stutus
	w.WriteHeader(204)
}
