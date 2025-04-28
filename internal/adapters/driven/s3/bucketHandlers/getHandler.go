package buckets

import (
	"Gononymous/internal/adapters/driven/s3/errs"
	"Gononymous/internal/adapters/driven/s3/structures"
	"Gononymous/internal/adapters/driven/s3/utils"
	"encoding/xml"
	"net/http"
	"os"
)

// Method for "GET" request
func GetHandler(w http.ResponseWriter, req *http.Request) {
	// Cheking for storage existense
	if _, err := os.Stat(StorageName + "/" + "buckets.csv"); err != nil {
		err := errs.Error{Code: 404, Message: "The storage is empty, put something before getting", Resource: req.URL.Path}
		err.WriteError(w)
		return
	}
	// Getting metadata of all buckets from buckets.csv
	buckets := utils.GetMetadata()
	buckets = buckets[1:]
	// Storing all data in one structure
	LAMB := structures.ListAllMyBucketsResult{AllBuckets: buckets}
	x, err := xml.MarshalIndent(LAMB, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Retriving the status
	w.WriteHeader(http.StatusOK)
	// Setting the header
	w.Header().Set("Content-Type", "application/xml")
	// Returning xml
	w.Write(x)
}
