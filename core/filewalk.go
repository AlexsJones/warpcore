package core

import (
	"os"
	"path/filepath"
)

//SearchFiles wraps FTW
func SearchFiles(filep string, token string) []string {

	fileSlice := make([]string, 0, 100)

	_ = filepath.Walk(filep, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == token {
			fileSlice = append(fileSlice, path)
		}
		return nil
	})

	return fileSlice
}
