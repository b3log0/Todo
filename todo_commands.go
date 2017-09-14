package main

import (
	"strings"
	"encoding/json"
	"strconv"
	"time"
	"bytes"
	"fmt"
	"path/filepath"
	"os"
	"bufio"
	"io"
	"os/exec"
)

func showTypes() {
	for index, domain := range getDomains() {
		printDomain(index, domain, strings.Compare(domain, current_domain) == 0)
	}
}

func removeByNumber(param int) {
	domain := getDomain(int64(param))
	if domain == current_domain {
		printError("Cannot delete the current todo list")
		return
	}
	delDomain(domain)
}

func listTasks() {
	taskList = getAllTasks(current_domain)
	for i := 0; i < len(taskList); i++ {
		temp := TaskDetail{}
		json.Unmarshal([]byte(taskList[i].taskDetail), &temp)
		printTask(i, temp)
	}
}

func cleanCurrentList() {
	for _, task := range taskList {
		temp := TaskDetail{}
		json.Unmarshal([]byte(task.taskDetail), &temp)
		if temp.State == 2 {
			delTask(current_domain, task.key)
		}
	}
}

func addNewTask(param string) {
	insertDomain(param)
}

func deleteTodoByNumber(params []string) {
	ids := getIdsFromParams(params)
	if len(ids) > 0 {
		for _, id := range ids {
			delTask(current_domain, taskList[id].key)
		}
	}
}

func doneByNumber(params []string) {
	task := TaskDetail{}
	ids := getIdsFromParams(params)
	if len(ids) > 0 {
		for _, id := range ids {
			json.Unmarshal([]byte(taskList[id].taskDetail), &task)
			task.State = 2
			setTask(current_domain, taskList[id].key, task.toJSONStr())
		}
	}
}

func doingByNumber(params []string) {
	task := TaskDetail{}
	ids := getIdsFromParams(params)
	if len(ids) > 0 {
		for _, id := range ids {
			json.Unmarshal([]byte(taskList[id].taskDetail), &task)
			task.State = 1
			setTask(current_domain, taskList[id].key, task.toJSONStr())
		}
	}
}

func commentByNumber(params []string) {
	task := TaskDetail{}
	id, _ := strconv.Atoi(params[0])
	comment := params[1:]
	json.Unmarshal([]byte(taskList[id].taskDetail), &task)
	task.Comment = strings.Join(comment, " ")
	setTask(current_domain, taskList[id].key, task.toJSONStr())
}

func notifyByNumber(params []string) {
	directory := os.Getenv("HOME")
    if directory == "" {
        directory = os.Getenv("USERPROFILE")
    }
    filename := filepath.Join(directory,TODO_CRON_FILE)
	w, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		printError("File Open Error!")
	}
	id, _ := strconv.Atoi(params[0])
	task := TaskDetail{}
	json.Unmarshal([]byte(getAllTasks(current_domain)[id].taskDetail), &task)
	cronexpr := ""
	for i:=1; i < len(params); i++{
		cronexpr=cronexpr + " " + params[i]
	}
	for i:=0; i<=5-len(params); i++{
		cronexpr = cronexpr + " *"
	}
	message := cronexpr + " export DISPLAY=:0.0;notify-send '" + task.Content + "' '" + task.Comment + "'\n"
	_, err = fmt.Fprintf(w, message)
	w.Close()
	exec.Command("crontab" , filename).Run()
}

func undoneByNumber(params []string) {
	task := TaskDetail{}
	ids := getIdsFromParams(params)
	if len(ids) > 0 {
		for _, id := range ids {
			json.Unmarshal([]byte(taskList[id].taskDetail), &task)
			task.State = 0
			setTask(current_domain, taskList[id].key, task.toJSONStr())
		}
	}
}

func listTasksByOrder(param int) {
	domain := getDomain(int64(param))
	setCurrentDomain(domain)
	taskList = getAllTasks(domain)
}

func exportAllTasksJSON(directory []string) {
	dir := directory[0]
	if dir == "" {
		dir = os.Getenv("HOME")
	}
	if dir == "" {
		dir = os.Getenv("USERPROFILE")
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	f, err := os.Create(filepath.Join(dir, TODO_EXPORT_JSON))
	defer f.Close()
	if err != nil {
		printError("Create File Error!")
		return
	}

	var buffer bytes.Buffer
	for _, domain := range getDomains() {
		taskList = getAllTasks(domain)
		for i := 0; i < len(taskList); i++ {
			buffer.WriteString(taskList[i].taskDetail + "\n")
		}
		buffer.WriteString("\n")
	}

	w, err := os.OpenFile(filepath.Join(dir, TODO_EXPORT_JSON), os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		printError("Open File Error")
		return
	}
	defer w.Close()
	fmt.Fprintf(w, buffer.String())
}

func exportAllTasksMD(directory []string) {
	dir := directory[0]
	if dir == "" {
		dir = os.Getenv("HOME")
	}
	if dir == "" {
		dir = os.Getenv("USERPROFILE")
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	f, err := os.Create(filepath.Join(dir, TODO_EXPORT_MD))
	defer f.Close()
	if err != nil {
		printError("Create File Error!")
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString("## Todo List" + "\n")
	for _, domain := range getDomains() {
		buffer.WriteString("### " + domain + "\n")
		taskList = getAllTasks(domain)
		for i := 0; i < len(taskList); i++ {
			temp := TaskDetail{}
			json.Unmarshal([]byte(taskList[i].taskDetail), &temp)
			if temp.State == 0 {
				buffer.WriteString("- [ ] " + temp.Content + "\n")
			} else if temp.State == 2 {
				buffer.WriteString("- [x] " + temp.Content + "\n")
			}
		}
		buffer.WriteString("\n")
	}

	w, err := os.OpenFile(filepath.Join(dir, TODO_EXPORT_MD), os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		printError("Open File Error")
		return
	}
	defer w.Close()
	fmt.Fprintf(w, buffer.String())
}

func importTasks(directory []string) {
	dir := directory[0]
	if dir == "" {
		dir = os.Getenv("HOME")
	}
	if dir == "" {
		dir = os.Getenv("USERPROFILE")
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	f, err := os.Open(filepath.Join(dir, TODO_EXPORT_JSON))
	if err != nil {
		printError("File Open Error")
		return
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		b, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				printError("File Read Error")
				return
			}
			break
		}
		line := string(b)
		if line != "" {
			temp := TaskDetail{}
			json.Unmarshal([]byte(line), &temp)
			if temp.Domain != "" && !DomainExists(getDomains(), temp.Domain) {
				insertDomain(temp.Domain)
			}
			setTask(temp.Domain, strconv.FormatInt(time.Now().UnixNano(), 10), line)
		}
	}
}

func addNewTodo(params []string) {
	setTask(current_domain, strconv.FormatInt(time.Now().UnixNano(), 10), newTaskDetail(strings.Join(params, " ")).toJSONStr())
}

func pushTasks(){
	for _, domain := range getDomains() {
		createBoard(domain)
	}
}

func pullTasks(){

}