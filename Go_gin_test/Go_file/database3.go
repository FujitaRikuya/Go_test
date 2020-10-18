package main

//import my originally package "dblib"
import (
	"fmt"

	"../Go_gin_test/dblib"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	myname := "Rikuya"
	Isaid := "Hi there:)"

	dblib.DbInit()
	dblib.DbInsert(myname, Isaid)

	fmt.Println("test.sqlite3")
	msg := dblib.DbGet()
	//fmt.Println(msg) 全部表示されてしまう

	//並べて表示してみる
	dblib.Disp(msg)
}
