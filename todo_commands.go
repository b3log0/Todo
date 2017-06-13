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
	current_mark = "_todo_"
	task_mark1 = " "
	task_mark2 = "*"
	done_mark1 = "[ ]"
	done_mark2 = "[*]"
)

func showTypes() {
	files, _ := ioutil.ReadDir(current_dir)
	n := 1
	for _,value := range files{
		if strings.HasSuffix(value.Name(),".todo") {
			fmt.Printf("%s %03d: %s\n", task_mark1, n, strings.TrimSpace(value.Name()))
			n++
		}
	}
}

func listTasks(filename string) error{
	todoType := strings.Replace(strings.Replace(filename,file_suffix,"",-1),current_mark,"",-1)
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
		if strings.HasPrefix(line, current_mark) {
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
		for index,value := range files{
			if index+1 == choose {
				if !strings.HasPrefix(value.Name(),current_mark) {
					os.Rename(getFilePathName(value.Name()),getFilePathName(current_mark+value.Name()))
					fileName = current_mark+value.Name()
				} else {
					fileName = value.Name()
				}
				fmt.Println(fileName)
			} else {
				name := strings.Replace(value.Name(),current_mark,"",-1)
				os.Rename(getFilePathName(value.Name()),getFilePathName(name))
			}
		}
		listTasks(fileName)
	}
}

func addNewTask(filename string) error {
	fmt.Println("create a new file :> "+filename)
	f,err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}
