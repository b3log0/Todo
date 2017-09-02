package main

import (
	"encoding/json"
	"time"
)

//redis hash key
//save as List
type Domain struct {
	Name string
	State bool
	Field string
}

//redis hash field
//save as Set
type Task struct {
	//注意大小写，大写表示public，小写表示private，json.Marshal要求public
	Content  string `json:"content"`
	State    bool 	`json:"state"`
	Comment  string `json:"comment,omitempty"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}



func (t Task) toJSON() []byte {
	task, _ := json.Marshal(&t)
	return task
}
