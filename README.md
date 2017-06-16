# Todo
a cli todo list 
**本项目改写自mattn大神的项目[todo](https://github.com/mattn/todo)**
mattn的todo简单好用，在使用过程中，我发现自己想要对不同的项目的todo清单进行分类，而原项目没有支持
于是自行改写成本项目，修改了命令行工具包（[cli](https://github.com/urfave/cli)比mattn使用的[commander](https://github.com/gonuts/commander)star数更多一些）

感谢大神的代码，让我省下了一半的时间（todo文件的读写程序完全照搬）

基于mattn的todo上，多了如下功能：
- new指令：创建一个新的todo类型
- remove指令：删除指定的数据类型

使用过程可能如下(图例中是在当前目录下执行的指令，所以有./，实际放到环境变量下的指令不需要加)：
```shell
todo #显示当前有何类型
```
![1](https://user-images.githubusercontent.com/2569600/27209005-cb59dfd0-527b-11e7-975c-b4beaab66452.png)
```shell
todo 1 # 选择序号为1的类型，此时当前工作的类型是1
```
![2](https://user-images.githubusercontent.com/2569600/27209007-cb5e8170-527b-11e7-83d1-560ba06247a1.png)
```shell
todo add something todo # 向类型1添加一个待办事项
```
![3](https://user-images.githubusercontent.com/2569600/27209006-cb5e5722-527b-11e7-8b95-4ae146b33bda.png)
```shell
todo add something todo 2 # 继续添加
todo add something todo 3 # 继续添加
```
![4](https://user-images.githubusercontent.com/2569600/27209008-cb6100da-527b-11e7-8bad-cf7c02a3687f.png)
```shell
todo done 1 # 完成了第一个事项
```
![5](https://user-images.githubusercontent.com/2569600/27209010-cb627956-527b-11e7-80b9-8961ef66fcbe.png)
```shell
todo delete 2 # 删除第二个待办事项
```
![6](https://user-images.githubusercontent.com/2569600/27209009-cb6219b6-527b-11e7-9a13-22ad3e71190b.png)
```shell
todo clean # 清除已完成的事项
```
![7](https://user-images.githubusercontent.com/2569600/27209011-cb865d8a-527b-11e7-9046-68d316cc951c.png)
```shell
todo 2 # 选择序号为2的类型，此时看到2类型下的待办事项
```
![8](https://user-images.githubusercontent.com/2569600/27209014-cb92656c-527b-11e7-944d-c181607573af.png)
```shell
todo remove 2 # 删除类型为2的待办事项
```
![9](https://user-images.githubusercontent.com/2569600/27209012-cb8e030a-527b-11e7-82f6-a88c099cdef4.png)
```shell
todo 1 # 再次选择类型为1的待办事项
```
```shell
... # other jobs
todo list # 显示当前工作的待办事项
```
![10](https://user-images.githubusercontent.com/2569600/27209013-cb8fbfe2-527b-11e7-9e7c-beb7625f3567.png)

欢迎各位使用 :tada:
