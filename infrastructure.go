package main

import (
	"github.com/go-redis/redis"
	"github.com/fatih/color"
	"sort"
	"time"
)

const (
	REDIS_KEY = "org.b3log.todo"
	TASK_MARK1 = " "
	TASK_MARK2 = "*"
	DONE_MARK1 = "\u2610" //on Mac
	DONE_MARK2 = "\u2611" //on Mac
	TODO_EXPORT_MD = "TodoList.md"
	TODO_EXPORT_JSON = "TodoList.json"
	TODO_CRON_FILE = "todo.cron"
	TODO_PROJ_HOME = "github.com/b3log/Todo"
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
		Domain:current_domain,
		Content:content,
		State:0,
		Comment:"",
		Created:time.Now().Format("2006-01-02 15:04:05"),
	}
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

func DomainExists(domains []string, domain string) bool{
	for _,value := range domains {
		if domain == value {
			return true
		}
	}
	return false
}