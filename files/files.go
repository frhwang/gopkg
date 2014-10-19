package files

import "os"

// IsNotExist returns true if a path isn't exist
func IsNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// IsExist returns true if a path is exist
func IsExist(path string) bool {
	return !IsNotExist(path)
}
