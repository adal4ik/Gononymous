package buckets

import (
	"encoding/xml"
	"net/http"
	"os"
	"time"
	"triple-s/internal/structures"
)

// Method for "PUT" request
func PutHandler(w http.ResponseWriter, req *http.Request) {
	// Getting bucket name
	dirName := req.PathValue("BucketName")

	// Checking for valide naming
	// passed := utils.ValidateDirName(dirName)
	// if !passed {
	// 	err := errs.Error{Code: 400, Message: "Bad Request for invalid names:" + dirName, Resource: req.URL.Path}
	// 	err.WriteError(w)
	// 	return
	// }

	// Cheking for duplicate names (existense of the bucket)
	// if _, err := os.Stat(StorageName + "/buckets.csv"); err == nil {
	// 	if utils.IsExists(dirName) {
	// 		err := errs.Error{Code: 409, Message: "Conflict for duplicate names:" + dirName, Resource: req.URL.Path}
	// 		err.WriteError(w)
	// 		return
	// 	}
	// }

	// Creating a new bucket
	os.MkdirAll(StorageName+"/"+dirName, os.ModePerm)
	// fmt.Println(StorageName + "/" + dirName)
	// if _, err := os.Stat(StorageName + "/" + dirName); os.IsNotExist(err) {
	// 	fmt.Println(StorageName+"/"+dirName, "does not exist")
	// } else {
	// 	fmt.Println("The provided directory named", StorageName+"/"+dirName, "exists")
	// }
	bucket := structures.Bucket{Name: dirName, CreatedTime: time.Now().Format(time.RFC1123), LastModifiedTime: time.Now().Format(time.RFC1123), Status: "Active"}

	// Preparing data about bucket to share
	x, err := xml.MarshalIndent(bucket, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Saving data about the bucket
	// utils.SaveToCsv(bucket)

	// Retrinig the result
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
