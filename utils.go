package main

import (
	"path/filepath"
)

func getFilePathName(filename string) string {
	return filepath.Join(current_dir,filename)
}