package main

import (
	"github.com/go-redis/redis"
	"github.com/fatih/color"
	"strconv"
	"sort"
	"time"
)

const (
	REDIS_KEY = "org.b3log.todo"
	CURRENT_KEY = "org.b3log.todo.current"
	TASK_MARK1 = " "
	TASK_MARK2 = "*"
	DONE_MARK1 = "\u2610" //on Mac
	DONE_MARK2 = "\u2611" //on Mac
	// done_mark1 = "[ ]" //on Windows
	// done_mark2 = "[*]" //on Windows
)

//redis client
var client *redis.Client
//all task list
var taskList []Task
//current task
var current_domain string

func initTodo(){
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		color.Red("Redis connection error!")
		return
	}
	current_domain = getCurrentDomain()
	taskList = getAllTasks(current_domain)
}

func newTaskDetail(content string) TaskDetail{
	return TaskDetail{
		Content:content,
		State:false,
		Comment:"",
		Created:time.Now().Format("2006-01-02 15:04:05"),
		Modified:time.Now().Format("2006-01-02 15:04:05"),
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

func buildTaskList(taskMap map[string]string) []Task{
	keySet := make([]string,0)
	for key,_ := range taskMap{
		keySet = append(keySet,key)
	}
	sort.Strings(keySet)
	taskList := make([]Task,len(keySet))
	for k,value := range keySet{
		taskList[k] = Task{
			key:value,
			taskDetail:taskMap[value],
		}
	}
	return taskList
}