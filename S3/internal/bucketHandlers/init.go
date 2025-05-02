package buckets

import (
	"os"
	"path/filepath"

	"triple-s/internal/utils"
)

// Init initializes the bucket handlers package
func Init() error {
	// Ensure buckets.csv exists
	csvPath := filepath.Join(utils.StorageBase, "buckets.csv")
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		headers := []string{"Name", "CreatedTime", "LastModifiedTime", "Status"}
		if err := utils.CreateCsv(csvPath, headers); err != nil {
			return err
		}
	}
	return nil
}
