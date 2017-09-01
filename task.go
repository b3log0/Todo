package main

import (
	"encoding/json"
	"time"
)

type Task struct {
	//注意大小写，大写表示public，小写表示private，json.Marshal要求public
	Domain   string `json:"domain"`
	Content  string `json:"content"`
	State    bool 	`json:"state"`
	Comment  string `json:"comment,omitempty"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

func newTask(domain string, content string) Task{
	return Task{
		Domain:domain,
		Content:content,
		State:false,
		Comment:"",
		Created:time.Now().Format("2006-01-02 15:04:05"),
		Modified:time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (t Task) toJSON() []byte {
	task, _ := json.Marshal(&t)
	return task
}
