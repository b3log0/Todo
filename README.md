# Todo
a cli todo list 
**本项目改写自mattn大神的项目[todo](https://github.com/mattn/todo)**
mattn的todo简单好用，在使用过程中，我发现自己想要对不同的项目的todo清单进行分类，而原项目没有支持
于是自行改写成本项目，修改了命令行工具包（[cli](github.com/urfave/cli)比mattn使用的[commander](github.com/gonuts/commander)star数更多一些）

基于mattn的todo上，多了如下功能：
```shell
# 新增任务类型
todo new test 
# 显示所有任务类型
todo
# 选择任务类型，参数为序号
todo 1
# 然后即可使用之前的add，clean，delete，done，list指令了
```
感谢大神的代码，让我省下了一半的时间

ps，生成的文件保存在编译时当前目录下的.todo文件夹中

现在已经可以满足一般需求，如果有时间，可能做如下更改：

- 在没有创建类型时，创建默认类型
- 将.todo文件夹生成在user目录下的固定位置
- 网络同步