package fileHandlers

import "triple-s/internal/utils"

var StorageName = "/data"

func Sync(newStorageName string) {
	StorageName = newStorageName
	utils.SyncDir(newStorageName)
}
