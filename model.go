package main

/**
 Status：
	Board	0 停用 1 启用
	Group	0 隐藏 1 展示
	Card	0 未完成 1 已完成 2 进行中
 */

type Board struct {
	Bid int
	Status int
	Id string
	Name string
	Desc string
}

type Group struct {
	Gid int
	Status int
	Id string
	Name string
	Bid int
}

type Card struct {
	Cid int
	Status int
	Id string
	Name string
	Desc string
	Gid int
}
