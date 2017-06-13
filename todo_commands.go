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
	curDir,_:=os.Getwd()
	files, _ := ioutil.ReadDir(curDir)
	n := 1
	for _,value := range files{
		if strings.HasSuffix(value.Name(),".todo") {
			fmt.Printf("%s %03d: %s\n", task_mark1, n, strings.TrimSpace(value.Name()))
			n++
		}
	}
}

func listTasks(filename string) error{
	todoType := strings.Replace(filename,file_suffix,"",-1)
	fmt.Println("=== "+todoType+" ===")
	f, err := os.Open(filename)//"./"+todoType+file_suffix
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
		if strings.HasPrefix(line, "-") {
			fmt.Printf("%s %03d: %s\n", done_mark2, n, strings.TrimSpace(string(line[1:])))
		} else {
			fmt.Printf("%s %03d: %s\n", done_mark1, n, strings.TrimSpace(line))
		}
		n++

	}
	return nil
}

func listTasksByOrder(order string) {
	choose,_:= strconv.Atoi(order)
	curDir,_ := os.Getwd()
	files, _ := ioutil.ReadDir(curDir)
	for index,value := range files{
		if index+1 == choose {
			listTasks(value.Name())
		}
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