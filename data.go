package main

import(
"database/sql"
"time"
_ "github.com/mattn/go-sqlite3"
)

const (
	BOARD_INSERT="insert into board(id,name,desc,created) values(?,?,?,?)"
	BOARD_UPDATE="update board set id=?,name=?,desc=?,created=? where bid=?"
	BOARD_DELETE="delete from board where bid=?"

	GROUP_INSERT="insert into group(id,name,created) values(?,?,?)"
	GROUP_UPDATE="update group set id=?,name=?,created=? where gid=?"
	GROUP_DELETE="delete from group where gid=?"

	CARD_INSERT="insert into card(id,name,desc,idList,created) values(?,?,?,?,?)"
	CARD_UPDATE="update card set id=?,name=?,desc=?,idList=?,created=? where cid=?"
	CARD_DELETE="delete from card where cid=?"
)

func main(){
	updateBoard()
}

func createBoard(board string,) {
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)

	stmt,err:=db.Prepare(BOARD_INSERT)
	checkErr(err)

	_,err=stmt.Exec("testid","testboard","test test test",time.Now().Format("2006-01-02 15:04:05"))
	checkErr(err)

}

func updateBoard(){
	db,err := sql.Open("sqlite3","./todo.db")
	checkErr(err)

	stmt,err:=db.Prepare(BOARD_UPDATE)
	checkErr(err)

	_,err=stmt.Exec("testid2","testboard2","test test test2",time.Now().Format("2006-01-02 15:04:05"),1)
	checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}