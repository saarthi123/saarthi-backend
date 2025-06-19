package utils

import "os"

// GetEnv returns the value of an environment variable or fallback if not set.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
