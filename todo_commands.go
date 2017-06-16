package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"bufio"
	"io"
	"os"
	"strconv"
	"github.com/fatih/color"
)

const (
	task_mark1 = " "
	task_mark2 = "*"
	done_mark0 = "-"
	done_mark1 = "\u2610" //on Mac
	done_mark2 = "\u2611" //on Mac
	// done_mark1 = "[ ]" //on Windows
	// done_mark2 = "[*]" //on Windows
)

func showTypes() {
	boldMagenta := color.New(color.FgMagenta).Add(color.Bold)
	files, _ := ioutil.ReadDir(current_dir)
	n := 1
	for _,value := range files{
		if strings.HasSuffix(value.Name(),todo_suffix) {
			color.Cyan("%s %03d: %s\n", task_mark1, n, strings.TrimSpace(strings.Replace(value.Name(),todo_suffix,"",-1)))
			n++
		}else if strings.HasSuffix(value.Name(),doing_suffix) {
			boldMagenta.Printf("%s %03d: %s\n", task_mark2, n, strings.TrimSpace(strings.Replace(value.Name(),doing_suffix,"",-1)))
			n++
		}
	}
}

func removeByNumber(num int) {
	files, _ := ioutil.ReadDir(current_dir)
	n := 1
	for _,value := range files{
		if strings.HasSuffix(value.Name(),todo_suffix) || strings.HasSuffix(value.Name(),doing_suffix) {
			if n == num {
				os.Remove(getFilePathName(value.Name()))
				break
			}
			n++
		}
	}
}

//此处的filename包含路径（在utils调用出加上了）
func listTasks(filename string,params []string) error{
	//如果进入其他目录执行指令，会出现此处显示为空的情况，这是由于当前目录获取错误导致
	//不知道是否需要进行处理，暂时不处理，因为正是生成的可执行文件应该是不会变路径的
	f, err := os.Open(filename)
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
		if strings.HasPrefix(line, done_mark0) {
			color.Green("%s %03d: %s\n", done_mark2, n, strings.TrimSpace(string(line[1:])))
		} else {
			color.Magenta("%s %03d: %s\n", done_mark1, n, strings.TrimSpace(line))
		}
		n++
	}
	return nil
}

//清除当前路径下已完成的列表
func cleanCurrentList(filename string,params []string) error{
	//创建一个临时文件，然后读取处理写入，删除原文件，新文件改名
	w, err := os.Create(filename + "_")
	if err != nil {
		return err
	}
	defer w.Close()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	br := bufio.NewReader(f)
	for {
		b, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		line := string(b)
		if !strings.HasPrefix(line, done_mark0) {
			_, err = fmt.Fprintf(w, "%s\n", line)
			if err != nil {
				return err
			}
		}
	}
	f.Close()
	w.Close()
	err = os.Remove(filename)
	if err != nil {
		return err
	}
	return os.Rename(filename+"_", filename)
}

func addNewTodo(filename string,params []string) error {
	w, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = fmt.Fprintf(w, " %s\n", strings.Join(params, " "))
	return err
}

func deleteTodoByNumber(filename string,params []string) error {
	ids := []int{}
	for _, arg := range params {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}
	w, err := os.Create(filename + "_")
	if err != nil {
		return err
	}
	defer w.Close()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
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
		match := false
		for _, id := range ids {
			if id == n {
				match = true
			}
		}
		if !match {
			_, err = fmt.Fprintf(w, "%s\n", string(b))
			if err != nil {
				return err
			}
		}
		n++
	}
	f.Close()
	w.Close()
	err = os.Remove(filename)
	if err != nil {
		return err
	}
	return os.Rename(filename+"_", filename)
}

func doneByNumber(filename string,params []string) error {
	ids := []int{}
	for _, arg := range params {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}
	w, err := os.Create(filename + "_")
	if err != nil {
		return err
	}
	defer w.Close()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
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
		match := false
		for _, id := range ids {
			if id == n {
				match = true
			}
		}
		line := strings.TrimSpace(string(b))
		if match && !strings.HasPrefix(line, "-") {
			_, err = fmt.Fprintf(w, "-%s\n", line)
			if err != nil {
				return err
			}
		} else {
			_, err = fmt.Fprintf(w, "%s\n", line)
			if err != nil {
				return err
			}
		}
		n++
	}
	f.Close()
	w.Close()
	err = os.Remove(filename)
	if err != nil {
		return err
	}
	return os.Rename(filename+"_", filename)
}

func undoneByNumber(filename string,params []string) error {
	ids := []int{}
	for _, arg := range params {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}
	w, err := os.Create(filename + "_")
	if err != nil {
		return err
	}
	defer w.Close()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
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
		match := false
		for _, id := range ids {
			if id == n {
				match = true
			}
		}
		line := strings.TrimSpace(string(b))
		if match && strings.HasPrefix(line, "-") {
			_, err = fmt.Fprintf(w, "%s\n", string(line[1:]))
			if err != nil {
				return err
			}
		} else {
			_, err = fmt.Fprintf(w, "%s\n", line)
			if err != nil {
				return err
			}
		}
		n++
	}
	f.Close()
	w.Close()
	err = os.Remove(filename)
	if err != nil {
		return err
	}
	return os.Rename(filename+"_", filename)
}

func listTasksByOrder(order string) {
	choose,err:= strconv.Atoi(order)
	if err != nil {
		color.Red("input a number")
	} else {
		files, _ := ioutil.ReadDir(current_dir)
		var filename string
		index := 0
		for _,value := range files{
			if strings.HasSuffix(value.Name(),todo_suffix) || strings.HasSuffix(value.Name(),doing_suffix){
				index = index + 1
			}
			if index == choose {
				if !strings.HasSuffix(value.Name(),doing_suffix) {
					filename = strings.Replace(value.Name(),todo_suffix,doing_suffix,-1)
					os.Rename(getFilePathName(value.Name()),getFilePathName(filename))
				} else {
					filename = value.Name()
				}
			} else {
				name := strings.Replace(value.Name(),doing_suffix,todo_suffix,-1)
				os.Rename(getFilePathName(value.Name()),getFilePathName(name))
			}
		}
		// todoType := strings.Replace(strings.Replace(filename,todo_suffix,"",-1),doing_suffix,"",-1)
		// color.Yellow("=== "+todoType+" ===")
		// listTasks(getFilePathName(filename),nil)
	}
}

func addNewTask(filename string) error {
	// type_name := strings.Replace(strings.Replace(filename,current_dir,"",-1),todo_suffix,"",-1)[1:]
	// fmt.Println("create a new type :> "+type_name)
	f,err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}
