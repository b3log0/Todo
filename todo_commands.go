package main

import (
	"strings"
	"encoding/json"
	"strconv"
	"time"
)

func showTypes() {
	for index, domain := range getDomains() {
		printDomain(index, domain, strings.Compare(domain,current_domain)==0)
	}
}

func removeByNumber(param int) {
	domain := getDomain(int64(param))
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
		if temp.State == true {
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
			task.State = true
			setTask(current_domain, taskList[id].key, task.toJSONStr())
		}
	}
}

func undoneByNumber(params []string) {
	task := TaskDetail{}
	ids := getIdsFromParams(params)
	if len(ids) > 0 {
		for _, id := range ids {
			json.Unmarshal([]byte(taskList[id].taskDetail), &task)
			task.State = false
			setTask(current_domain, taskList[id].key, task.toJSONStr())
		}
	}
}

func listTasksByOrder(param int) {
	domain := getDomain(int64(param))
	setCurrentDomain(domain)
	taskList = getAllTasks(domain)
	//for index, task := range taskList {
	//	temp := TaskDetail{}
	//	json.Unmarshal([]byte(task.taskDetail), &temp)
	//	printTask(index, temp)
	//}
}

func addNewTodo(params []string) {
	setTask(current_domain, strconv.FormatInt(time.Now().Unix(), 10), newTaskDetail(strings.Join(params, " ")).toJSONStr())
}