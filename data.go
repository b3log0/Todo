package main

import(
	"database/sql"
	"time"
_ 	"github.com/mattn/go-sqlite3"
)

const (
	BOARD_INSERT="insert into board(status,id,name,desc,created) values(?,?,?,?)"
	BOARD_UPDATE="update board set id=?,name=?,desc=?,created=? where bid=?"
	BOARD_DELETE="delete from board where bid=?"

	GROUP_INSERT="insert into group(status,id,name,bid,created) values(?,?,?,?)"
	GROUP_UPDATE="update group set id=?,name=?,bid=?,created=? where gid=?"
	GROUP_DELETE="delete from group where gid=?"

	CARD_INSERT="insert into card(status,id,name,desc,gid,created) values(?,?,?,?,?)"
	CARD_UPDATE="update card set id=?,name=?,desc=?,gid=?,created=? where cid=?"
	CARD_DELETE="delete from card where cid=?"

	DATE_FORMAT="2006-01-02 15:04:05"
)

func createBoard(board Board) {
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(BOARD_INSERT)
	checkErr(err)
	_,err=stmt.Exec(board.Status,board.Id,board.Name,board.Desc,time.Now().Format(DATE_FORMAT))
	checkErr(err)
}

func updateBoard(board Board){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(BOARD_UPDATE)
	checkErr(err)
	_,err=stmt.Exec(board.Status,board.Id,board.Name,board.Desc,time.Now().Format(DATE_FORMAT),board.Bid)
	checkErr(err)
}

func deleteBoard(id int){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(BOARD_DELETE)
	checkErr(err)
	_,err=stmt.Exec(id)
	checkErr(err)
}

func createGroup(group Group){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(GROUP_INSERT)
	checkErr(err)
	_,err=stmt.Exec(group.Status,group.Id,group.Name,time.Now().Format(DATE_FORMAT))
	checkErr(err)
}

func updateGroup(group Group){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(GROUP_UPDATE)
	checkErr(err)
	_,err=stmt.Exec(group.Status,group.Id,group.Name,time.Now().Format(DATE_FORMAT),group.Gid)
	checkErr(err)
}

func deleteGroup(id int){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(GROUP_DELETE)
	checkErr(err)
	_,err=stmt.Exec(id)
	checkErr(err)
}

func createCard(card Card){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(CARD_INSERT)
	checkErr(err)
	_,err=stmt.Exec(card.Status,card.Id,card.Name,card.Desc,time.Now().Format(DATE_FORMAT))
	checkErr(err)
}

func updateCard(card Card){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(CARD_UPDATE)
	checkErr(err)
	_,err=stmt.Exec(card.Status,card.Id,card.Name,card.Desc,time.Now().Format(DATE_FORMAT),card.Cid)
	checkErr(err)
}

func deleteCard(id int){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)
	stmt,err:=db.Prepare(CARD_DELETE)
	checkErr(err)
	_,err=stmt.Exec(id)
	checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}