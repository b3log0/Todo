package main

import (
	"path/filepath"
	"io/ioutil"
	"strings"
	"os/exec"
	"fmt"
	"time"
	"strconv"
	"github.com/fatih/color"
	//"github.com/jasonlvhit/gocron"
)

// func newCommand(number int64,order string) Command{
// 	return Command{
// 		Number:number,
// 		Order:order,
// 	}
// }

func newTask(content string) Task{
	return Task{
		Content:content,
		State:false,
		Comment:"",
		Created:time.Now().Format("2006-01-02 15:04:05"),
		Modified:time.Now().Format("2006-01-02 15:04:05"),
	}
}

//参数的filename不包含路径
func getFilePathName(filename string) string {
	return filepath.Join(current_dir,filename)
}

//参数的filename包含路径
func editDoingFunc(doingFunc func([]string),params []string) error {
	files, _ := ioutil.ReadDir(current_dir)
	for _,value := range files{
		if strings.HasSuffix(value.Name(),DOING_SUFFIX) {
			doingFunc(params)
		}
	}
	return nil
}

func printTask(num int,task Task){
	if task.State == true {
		color.Green("%s %03d: %s\n", DONE_MARK2, num, strings.TrimSpace(task.Content))
	}else{
		color.Magenta("%s %03d: %s\n", DONE_MARK1, num, strings.TrimSpace(task.Content))
	}
}

func printDomain(num int,domain string,current bool){
	boldMagenta := color.New(color.FgMagenta).Add(color.Bold)
	if current==true {
		boldMagenta.Printf("%s %03d: %s\n", TASK_MARK2, num, strings.TrimSpace(domain))
	}else{
		color.Cyan("%s %03d: %s\n", TASK_MARK1, num, strings.TrimSpace(domain))
	}
}

func getFirstIdFromParams(params []string) int{
	id,err := strconv.Atoi(params[0])
	if err == nil {
		return id
	}else{
		return -1
	}
}

func getIdsFromParams(params []string) []int{
	ids := []int{}
	for _, arg := range params {
		id, err := strconv.Atoi(arg)
		if err == nil {
			ids = append(ids, id)
		}
		
	}
	return ids
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