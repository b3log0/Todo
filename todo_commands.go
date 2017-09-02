package main

import (
	// "fmt"
	// "strings"
	// "io/ioutil"
	// "bufio"
	// "io"
	// "os"
	// "strconv"
	// "github.com/fatih/color"
	"encoding/json"
)

type command struct{
	order string
	number int64
	second int
	minitue int
	hour int
	day int
	month int
}

type todoCommand interface{
	showTypes() []string
	removeByNumber(num int64)
	listTasks() []string
	cleanCurrentList()
	addNewTask(domain string) bool
	deleteTodoByNumber(num int64)
	doneByNumber(num int64)
	undoneByNumber(num int64)
	listTasksByOrder(num int64)
	addNewTodo() error
}

func (c command) showTypes() []string{
	return getDomains()
}

func (c command) removeByNumber(num int64){
	domain := getDomain(num)
	delDomain(domain)
}

func (c command) listTasks() []string{
	current :=getCurrentDomain()
	return getTasks(current)
}

func (c command) cleanCurrentList() {
	current :=getCurrentDomain() //应该可以不用每次都获取一遍
	for _,task := range getTasks(current) {
		temp := Task{}
		json.Unmarshal([]byte(task),&temp)
		if temp.State == true {
			delTask(current,task)
		}
	}
}

func (c command) addNewTask(domain string) bool{
	return insertDomain(domain)
}

func (c command) deleteTodoByNumber(num int64) {
	current := getCurrentDomain()
	task := getTask(current,num)
	delTask(current,task)
}

func (c command) doneByNumber(num int64) {
	task := Task{}
	current := getCurrentDomain()
	json.Unmarshal([]byte(getTask(current,num)),&task)
	task.State = false
	setTask(current,task.toJSONStr())
}

func (c command) undoneByNumber(num int64) {
	task := Task{}
	current := getCurrentDomain()
	json.Unmarshal([]byte(getTask(current,num)),&task)
	task.State = true
	setTask(current,task.toJSONStr())
}

func (c command) listTasksByOrder(num int64) []string{
	domain := getDomain(num)
	return getTasks(domain)
}

func (c command) addNewTodo(content string) {
	current := getCurrentDomain()
	setTask(current,newTask(content).toJSONStr())
}