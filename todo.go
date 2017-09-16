package main

import (
    "github.com/urfave/cli"
    "os"
    "sort"
    "strconv"

)



func main() {
    //考虑此处进行初始化redis
    initTodo()

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
                listTasks()
                return nil
            },
        },
        {
            Name: "trello",
            Aliases: []string{"tr"},
            Action: func(c *cli.Context) error{
                trelloTasks()
                return nil
            },
        },
        {
            Name: "push",
            Aliases: []string{"ps"},
            Action: func(c *cli.Context) error{
                pushTasks()
                return nil
            },
        },
        {
            Name: "pull",
            Aliases: []string{"pl"},
            Action: func(c *cli.Context) error{
                pullTasks()
                return nil
            },
        },
        {
            Name: "export",
            Aliases: []string{"ex"},
            Action: func(c *cli.Context) error{
                if c.NArg() > 0{
                    exportAllTasksJSON(c.Args())
                }else{
                    return cli.NewExitError("need a file path",103)
                }
                return nil
            },
        },
        {
            Name: "generate",
            Aliases: []string{"ge"},
            Action: func(c *cli.Context) error{
                if c.NArg() > 0{
                    exportAllTasksMD(c.Args())
                }else{
                    return cli.NewExitError("need a file path",103)
                }
                return nil
            },
        },
        {
            Name: "import",
            Aliases: []string{"imp"},
            Action: func(c *cli.Context) error{
                if c.NArg() > 0{
                    importTasks(c.Args())
                }else{
                    return cli.NewExitError("need a file path",103)
                }
                return nil
            },
        },
        {
            Name: "add",
            Aliases: []string{"a"},
            Action: func(c *cli.Context) error{
                if c.NArg() > 0{
                    addNewTodo(c.Args())
                    listTasks()
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
                cleanCurrentList()
                listTasks()
                return nil
            },
        },
        {
            Name:"delete",
            Aliases:[]string{"dd"},
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    deleteTodoByNumber(c.Args())
                    listTasks()
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
                    doneByNumber(c.Args())
                    listTasks()
                }else{
                    return cli.NewExitError("need input a task number",104)
                }
                return nil
            },
        },
        {
            Name:"doing",
            Aliases:[]string{"doing"},
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    doingByNumber(c.Args())
                    listTasks()
                }else{
                    return cli.NewExitError("need input a task number",104)
                }
                return nil
            },
        },
        {
            Name:"comment",
            Aliases:[]string{"comment"},
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    doneByNumber(c.Args())
                    listTasks()
                }else{
                    return cli.NewExitError("need input a task number",104)
                }
                return nil
            },
        },
        {
            Name:"notify",
            Aliases:[]string{"n"},
            Action: func(c *cli.Context) error {
                if c.NArg() > 0{
                    notifyByNumber(c.Args())
                    listTasks()
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
                    addNewTask(task_name)
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
                    undoneByNumber(c.Args())
                    listTasks()
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
            arg,err := strconv.Atoi(c.Args()[0])
            if err!= nil{
                return cli.NewExitError("need input a task number",104)
            }
            listTasksByOrder(arg)
        }
        showTypes()
        return nil
    }
    sort.Sort(cli.FlagsByName(app.Flags))
    sort.Sort(cli.CommandsByName(app.Commands))
    app.Run(os.Args)
}
