package main

type Board struct {
	boardId string
	name string
	listsIds []List
}

type List struct {
	listId string
	name string
	cardIds []Card
}

type Card struct {
	cardId string
	title string
	description string
}

type BoardResp struct {
	Id string
	Name string
	Desc string 
}

type CardResp struct {
	Id string
	Name string
	Desc string
	IdList string
}