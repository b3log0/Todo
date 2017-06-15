# Todo
a cli todo list 
**本项目改写自mattn大神的项目[todo](https://github.com/mattn/todo)**
mattn的todo简单好用，在使用过程中，我发现自己想要对不同的项目的todo清单进行分类，而原项目没有支持
于是自行改写成本项目，修改了命令行工具包（[cli](https://github.com/urfave/cli)比mattn使用的[commander](https://github.com/gonuts/commander)star数更多一些）

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

欢迎各位使用 :tada:
