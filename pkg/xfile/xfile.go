package xfile

import (
	"fmt"
	"os"
	"path/filepath"
)

// IsDir is used to judge whether the specific path is a valid directory
func IsDir(path string) bool {
	var (
		fileInfo os.FileInfo
		err      error
	)
	fileInfo, err = os.Stat(path)

	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// IsFile is used to judge whether the specific path is a valid file path
func IsFile(path string) bool {
	var (
		fileInfo os.FileInfo
		err      error
	)
	fileInfo, err = os.Stat(path)

	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func CreateAndOpen(file string) (*os.File, error) {
	dir := filepath.Dir(file)

	if IsFile(dir) {
		return nil, fmt.Errorf("invalid file paht: %s", file)
	}

	if !IsDir(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	return os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
}
