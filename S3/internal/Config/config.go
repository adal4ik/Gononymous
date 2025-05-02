package config

import "os"

type Config struct {
	Port       string
	StorageDir string
	Debug      bool
}

func Load() Config {
	return Config{
		Port:       os.Getenv("PORT"),
		StorageDir: os.Getenv("STORAGE_DIR"),
		Debug:      os.Getenv("DEBUG") == "true",
	}
}
