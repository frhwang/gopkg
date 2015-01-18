package files

import "os"

// IsNotExist returns true if a path doesn't exist
func IsNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// IsExist returns true if a path exists
func IsExist(path string) bool {
	return !IsNotExist(path)
}
