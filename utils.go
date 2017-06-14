package main

import (
	"path/filepath"
	"io/ioutil"
	"strings"
)

//参数的filename不包含路径
func getFilePathName(filename string) string {
	return filepath.Join(current_dir,filename)
}

//参数的filename包含路径
func editDoingFunc(doingFunc func(string,[]string) error,params []string) error {
	files, _ := ioutil.ReadDir(current_dir)
	for _,value := range files{
		if strings.HasSuffix(value.Name(),doing_suffix) {
			doingFunc(getFilePathName(value.Name()),params)
		}
	}
	return nil
}