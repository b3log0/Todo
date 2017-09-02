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

type command struct{
	order string
	number int
	second int
	minitue int
	hour int
	day int
	month int
}

type todoCommand interface{
	showTypes()
	removeByNumber()
	listTasks() error
	cleanCurrentList() error
	addNewTodo() error
	deleteTodoByNumber() error
	doneByNumber() error
	undoneByNumber() error
	listTasksByOrder()
	addNewTask() error
}

func (c command) showTypes(){

}

func (c command) removeByNumber(){

}

func (c command) listTasks(){

}

func (c command) cleanCurrentList() {
	
}

func (c command) addNewTodo() {

}

func (c command) deleteTodoByNumber() {

}

func (c command) doneByNumber() {

}

func (c command) undoneByNumber() {

}

func (c command) listTasksByOrder() {

}

func (c command) addNewTask() {
	
}