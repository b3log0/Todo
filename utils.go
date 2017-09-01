package main

import (
	"path/filepath"
	"io/ioutil"
	"strings"
	"os/exec"
	"fmt"
	//"github.com/jasonlvhit/gocron"
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

func sendNotify(title string, message string,imagePath string) error {
	args := []string{}
	if imagePath != "" {
		args = append(args, "-i", imagePath)
	}
	args = append(args, title)
	args = append(args, message)
	cmd := exec.Command("notify-send", args...)
	if cmd == nil {
		return fmt.Errorf("No command")
	}
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error running command: %s", err)
	}
	return nil
}

//gocron.Every(5).Seconds().Do(sendNotify, "hello", "This notify message is from zephyr","/home/zephyr/Documents/GithubSpace/Todo/pics/icon.png")
//<- gocron.Start()