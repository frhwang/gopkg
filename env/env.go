package env

import "os"

// GetOrDefault should return a value from environment variable,
// but return a default value if variable isn't exist
func GetOrDefault(name, defaultValue string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return defaultValue
}
