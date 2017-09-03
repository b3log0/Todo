package main

import (
	// "fmt"
	"strings"
	// "io/ioutil"
	// "bufio"
	// "io"
	// "os"
	// "strconv"
	// "github.com/fatih/color"
	"encoding/json"
	//"fmt"
	"strconv"
)

type Command struct {
	Order   []string

	second  int
	minitue int
	hour    int
	day     int
	month   int
	year    int
}

// type todoCommand interface{
// 	showTypes()
// 	removeByNumber(num int64)
// 	listTasks()
// 	cleanCurrentList()
// 	addNewTask(domain string)
// 	deleteTodoByNumber(num int64)
// 	doneByNumber(num int64)
// 	undoneByNumber(num int64)
// 	listTasksByOrder(num int64)
// 	addNewTodo()
// }

func showTypes() {
	current := getCurrentDomain()
	for index, domain := range getDomains() {
		printDomain(index, domain, domain == current)
	}
}

func removeByNumber(param int) {
	domain := getDomain(int64(param))
	delDomain(domain)
}

func listTasks() {
	current := getCurrentDomain()
	tasks := getAllTasks(current)
	for i := 0; i < len(tasks); i++ {
		temp := Task{}
		json.Unmarshal([]byte(tasks[strconv.Itoa(i)]), &temp)
		printTask(i, temp)
	}
}

func cleanCurrentList() {
	current := getCurrentDomain() //应该可以不用每次都获取一遍
	for index, task := range getAllTasks(current) {
		temp := Task{}
		json.Unmarshal([]byte(task), &temp)
		if temp.State == true {
			delTask(current, index)
		}
	}
}

func addNewTask(param string) {
	insertDomain(param)
}

func deleteTodoByNumber(params []string) {
	ids := getIdsFromParams(params)
	current := getCurrentDomain()
	if len(ids) > 0 {
		for _, id := range ids {
			delTask(current, strconv.Itoa(id))
		}
	}
}

func doneByNumber(params []string) {
	task := Task{}
	current := getCurrentDomain()
	ids := getIdsFromParams(params)
	if len(ids) > 0 {
		for _, id := range ids {
			json.Unmarshal([]byte(getTask(current, int64(id))), &task)
			task.State = true
			setTask(current, int64(id), task.toJSONStr())
		}
	}
}

func undoneByNumber(params []string) {
	task := Task{}
	current := getCurrentDomain()
	ids := getIdsFromParams(params)
	if len(ids) > 0 {
		for _, id := range ids {
			json.Unmarshal([]byte(getTask(current, int64(id))), &task)
			task.State = false
			setTask(current, int64(id), task.toJSONStr())
		}
	}
}

func listTasksByOrder(param int) {
	domain := getDomain(int64(param))
	setCurrentDomain(domain)
	for index, task := range getAllTasks(domain) {
		temp := Task{}
		json.Unmarshal([]byte(task), &temp)
		i,_:=strconv.Atoi(index)
		printTask(i, temp)
	}
}

func addNewTodo(params []string) {
	current := getCurrentDomain()
	len := getTaskCount(current)
	setTask(current, len, newTask(strings.Join(params, " ")).toJSONStr())
}