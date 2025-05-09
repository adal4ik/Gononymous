package utils

import (
	"regexp"
	"testing"
)

func TestUUID(t *testing.T) {
	// Call the UUID function
	uuid := UUID()

	// Check if UUID is empty
	if uuid == "" {
		t.Error("UUID should not be empty")
	}

	// Use a regular expression to check the format of the UUID
	// The format is expected to be "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	// where x is a hexadecimal digit (0-9, a-f)
	re := regexp.MustCompile(`^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$`)

	if !re.MatchString(uuid) {
		t.Errorf("UUID format is invalid: %s", uuid)
	}
}
