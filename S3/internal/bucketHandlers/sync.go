package buckets

import "triple-s/internal/utils"

// Storage directory name for buckets package
var StorageName = "/data"

// Syncronization function
func Sync(newStorageName string) {
	StorageName = newStorageName
	utils.SyncDir(newStorageName)
}
