package files

import "os"

func IsNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func IsExist(path string) bool {
	return !IsNotExist(path)
}
