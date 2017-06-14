package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"bufio"
	"io"
	"os"
	"strconv"
)

const (
	task_mark1 = " "
	task_mark2 = "*"
	done_mark1 = "[ ]"
	done_mark2 = "[*]"
)

func showTypes() {
	files, _ := ioutil.ReadDir(current_dir)
	n := 1
	for _,value := range files{
		if strings.HasSuffix(value.Name(),todo_suffix) {
			fmt.Printf("%s %03d: %s\n", task_mark1, n, strings.TrimSpace(strings.Replace(value.Name(),todo_suffix,"",-1)))
			n++
		}else if strings.HasSuffix(value.Name(),doing_suffix) {
			fmt.Printf("%s %03d: %s\n", task_mark2, n, strings.TrimSpace(strings.Replace(value.Name(),doing_suffix,"",-1)))
			n++
		}
	}
}

func listTasks(filename string) error{
	todoType := strings.Replace(strings.Replace(filename,todo_suffix,"",-1),doing_suffix,"",-1)
	//如果进入其他目录执行指令，会出现此处显示为空的情况，这是由于当前目录获取错误导致
	//不知道是否需要进行处理，暂时不处理，因为正是生成的可执行文件应该是不会变路径的
	fmt.Println("=== "+todoType+" ===")
	f, err := os.Open(getFilePathName(filename))
	if err != nil {
		return err
	}
	defer f.Close()
	br := bufio.NewReader(f)
	n := 1
	for {
		b, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		line := string(b)
		if strings.HasPrefix(line, "_") {
			fmt.Printf("%s %03d: %s\n", done_mark2, n, strings.TrimSpace(string(line[1:])))
		} else {
			fmt.Printf("%s %03d: %s\n", done_mark1, n, strings.TrimSpace(line))
		}
		n++

	}
	return nil
}

func listTasksByOrder(order string) {
	choose,err:= strconv.Atoi(order)
	if err != nil {
		fmt.Println("input a number")
	} else {
		files, _ := ioutil.ReadDir(current_dir)
		var fileName string
		index := 0
		for _,value := range files{
			if strings.HasSuffix(value.Name(),todo_suffix) || strings.HasSuffix(value.Name(),doing_suffix){
				index = index + 1
			}
			if index == choose {
				if !strings.HasSuffix(value.Name(),doing_suffix) {
					fileName = strings.Replace(value.Name(),todo_suffix,doing_suffix,-1)
					os.Rename(getFilePathName(value.Name()),getFilePathName(fileName))
				} else {
					fileName = value.Name()
				}
			} else {
				name := strings.Replace(value.Name(),doing_suffix,todo_suffix,-1)
				os.Rename(getFilePathName(value.Name()),getFilePathName(name))
			}
		}
		listTasks(fileName)
	}
}

func addNewTask(filename string) error {
	type_name := strings.Replace(strings.Replace(filename,current_dir,"",-1),todo_suffix,"",-1)[1:]
	fmt.Println("create a new type :> "+type_name)
	f,err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}
