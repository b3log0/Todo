package main

import (
	"encoding/json"
	// "time"
)

//redis hash key
//save as List
// type Domain struct {
// 	Name string
// }

type Task struct{
	key string
	taskDetail string
}


//redis hash field
//save as Set
type TaskDetail struct {
	//注意大小写，大写表示public，小写表示private，json.Marshal要求public
	//State: 0 todo 1 doing 2 done
	Domain string `json:"domain"`
	Content  string `json:"content"`
	State    int 	`json:"state"`
	Comment  string `json:"comment,omitempty"`
	Created  string `json:"created"`
	Notify string `json:"notify,omitempty"`
}



func (t TaskDetail) toJSONStr() string {
	task, _ := json.Marshal(&t)
	return string(task)
}
