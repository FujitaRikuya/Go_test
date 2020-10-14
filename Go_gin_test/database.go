package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

//MLog ; type for Messagelog on Chat
type MLog struct {
	gorm.Model
	Name string
	Text string
}

func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInit)")
	}

	db.Create(&MLog{})
	defer db.Close()
}

func dbInsert(name string, text string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInsert)")
	}

	db.Create(&MLog{Name: name, Text: text})
	defer db.Close()
}

func dbGet() []MLog {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbGet")
	}

	var mlog []MLog
	db.Order("created_at desc").Find(&mlog)
	db.Close()

	return mlog
}

func main() {
	myname := "Rikuya"
	Isaid := "Hi there:)"
	defer fmt.Println(myname, Isaid)

	dbInit()
	dbInsert(myname, Isaid)

	msg := dbGet()
	fmt.Println(msg)

}
