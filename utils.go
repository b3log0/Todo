package main

import (
	"strings"
	"os/exec"
	"fmt"
	"github.com/fatih/color"
	//"github.com/jasonlvhit/gocron"
)

// func newCommand(number int64,order string) Command{
// 	return Command{
// 		Number:number,
// 		Order:order,
// 	}
// }




func printTask(num int,taskDetail TaskDetail){
	if taskDetail.State == true {
		color.Green("%s %03d: %s\n", DONE_MARK2, num, strings.TrimSpace(taskDetail.Content))
	}else{
		color.Magenta("%s %03d: %s\n", DONE_MARK1, num, strings.TrimSpace(taskDetail.Content))
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