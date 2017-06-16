package main

import (
    "github.com/urfave/cli"
    "os"
    "sort"
    "path/filepath"
    "strconv"
)

const (
    doing_suffix = ".doing"
    todo_suffix = ".todo"
    default_type_name = "main"
)

var current_dir string
    
func main() {
    current_dir = os.Getenv("HOME")
    if current_dir == "" {
        current_dir = os.Getenv("USERPROFILE")
    }
    current_dir = filepath.Join(current_dir,todo_suffix)
    _,err := os.Stat(current_dir)
    if os.IsNotExist(err) {
        os.Mkdir(current_dir,os.ModePerm)
        addNewTask(filepath.Join(current_dir, default_type_name + doing_suffix))
    }
    var todoType string
    app := cli.NewApp()
    app.Name = "ToDo"
    app.Usage = "A simple todo list"
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name: "type,t",
            Value: "main",
            Usage: "work space of todo list",
            Destination: &todoType,
        },
    }
    app.Commands = []cli.Command{
        {
            Name: "list",
            Aliases: []string{"l"},
            Action: func(c *cli.Context) error{
                editDoingFunc(listTasks,nil)
                return nil
            },
        },
        {
            Name: "add",
            Aliases: []string{"a"},
            Action: func(c *cli.Context) error{
                if c.NArg() > 0{
                    editDoingFunc(addNewTodo,c.Args())
                    editDoingFunc(listTasks,nil)
                }else{
                    return cli.NewExitError("need input a task",103)
                }
                return nil
            },
        },
        {
            Name: "clean",
            Aliases:[]string{"c"},
            Action: func(c *cli.Context) error {
                editDoingFunc(cleanCurrentList,nil)
                editDoingFunc(listTasks,nil)
                return nil
            },
        },
        {
            Name:"delete",
            Aliases:[]string{"dd"},
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    editDoingFunc(deleteTodoByNumber,c.Args())
                    editDoingFunc(listTasks,nil)
                }else{
                    return cli.NewExitError("need input a task number",104)
                }
                return nil
            },
        },
        {
            Name:"done",
            Aliases:[]string{"d"},
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    editDoingFunc(doneByNumber,c.Args())
                    editDoingFunc(listTasks,nil)
                }else{
                    return cli.NewExitError("need input a task number",104)
                }
                return nil
            },
        },
        {
            Name:"new",
            Aliases:[]string{"n"},
            Usage: "create a new file for task",
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    task_name := c.Args()[0]
                    errMsg := addNewTask(filepath.Join(current_dir, task_name + todo_suffix))
                    if errMsg != nil {
                        return cli.NewExitError("cannot create a new file for the list",101)
                    }
                }else{
                    return cli.NewExitError("create new todo list error, a name is required",100)
                }
                showTypes()
                return nil
            },
        },
        {
            Name:"undone",
            Aliases:[]string{"u"},
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    editDoingFunc(undoneByNumber,c.Args())
                    editDoingFunc(listTasks,nil)
                }else{
                    return cli.NewExitError("need input a task number",104)
                }
                return nil
            },
        },
        {
            Name:"remove",
            Aliases:[]string{"r"},
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    arg,err := strconv.Atoi(c.Args()[0])
                    if err!= nil{
                        return cli.NewExitError("need input a task number",104)
                    }
                    removeByNumber(arg)
                    showTypes()
                }else{
                    return cli.NewExitError("need input a task number",104)
                }
                return nil
            },
        },
    }

    //定义一级指令行为，有参数时显示指定序号的todo清单内容，否则显示所有todo清单名称
    app.Action = func(c *cli.Context) error {
        if c.NArg() > 0{
            listTasksByOrder(c.Args()[0])
        }
        showTypes()
        return nil
    }
    sort.Sort(cli.FlagsByName(app.Flags))
    sort.Sort(cli.CommandsByName(app.Commands))
    app.Run(os.Args)
}
