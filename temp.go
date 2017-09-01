package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	task1 := newTask("main", "test test test3")
	task2 := newTask("fresh", "test test test1")
	insertTask("main", string(task1.toJSON()))
	insertTask("fresh", string(task2.toJSON()))
	task3 := Task{}
	task4 := Task{}
	json.Unmarshal([]byte(getTaskByDomain("main")), &task3)
	json.Unmarshal([]byte(getTaskByDomain("fresh")), &task4)
	fmt.Println(task3)
	fmt.Println(task4)
}
