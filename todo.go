package main

import (
    "fmt"
    "github.com/urfave/cli"
    "os"
    "sort"
    "path/filepath"
)

const (
    file_suffix = ".todo"
)
    
func main() {
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
            Name: "add",
            Aliases: []string{"a"},
            Action: func(c *cli.Context) error{
                fmt.Println("complete!")
                return nil
            },
        },
        {
            Name: "clean",
            Aliases:[]string{"a"},
            Action: func(c *cli.Context) error {
                fmt.Println("add!")
                return nil
            },
        },
        {
            Name:"delete",
            Aliases:[]string{"dd"},
            Action: func(c *cli.Context) error {
                fmt.Println("add!")
                return nil
            },
        },
        {
            Name:"done",
            Aliases:[]string{"d"},
            Action: func(c *cli.Context) error {
                fmt.Println("add!")
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
                    curDir,_:=os.Getwd()
                    errMsg := addNewTask(filepath.Join(curDir, task_name + file_suffix))
                    if errMsg != nil {
                        return cli.NewExitError("cannot create a new file for the list",101)
                    }
                }else{
                    return cli.NewExitError("create new todo list error, a name is required",100)
                }
                return nil
            },
        },
        {
            Name:"undone",
            Aliases:[]string{"u"},
            Action: func(c *cli.Context) error {
                fmt.Println("add!")
                return nil
            },
        },
    }

    //定义一级指令行为，有参数时显示指定序号的todo清单内容，否则显示所有todo清单名称
    app.Action = func(c *cli.Context) error {
        if c.NArg() > 0{
            listTasksByOrder(c.Args()[0])
        }else{
            showTypes()
        }
        
        return nil
    }
    sort.Sort(cli.FlagsByName(app.Flags))
    sort.Sort(cli.CommandsByName(app.Commands))
    app.Run(os.Args)
}
