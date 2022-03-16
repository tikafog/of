package utils

import (
	"os"
)

// FileExists check if the file with the given path exits.
func FileExists(filename string) bool {
	fi, err := os.Lstat(filename)
	if fi != nil || (err != nil && !os.IsNotExist(err)) {
		return true
	}
	return false
}

// FileNotExists check if the file with the given path not exits.
func FileNotExists(filename string) bool {
	fi, err := os.Lstat(filename)
	if fi == nil || err != nil && os.IsNotExist(err) {
		return true
	}
	return false
}
