package files

import "os"

// IsNotExist return true if a path is exist
func IsNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// IsExist return true if a path isn't exist
func IsExist(path string) bool {
	return !IsNotExist(path)
}
