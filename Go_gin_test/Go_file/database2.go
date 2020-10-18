package main

import (
	"fmt"

	"gorm.io/gorm"

	_ "gorm.io/driver/sqlite"
	/*
		"github.com/jinzhu/gorm"

		_ "github.com/mattn/go-sqlite3"
	*/)

//MLog ; type for Messagelog on Chat
type MLog struct {
	gorm.Model
	Name string
	Text string
}

func display([]msg Mlog){
  i := 1
  for _, text := range msg {
    fmt.Println(i , "; "  text.Text)
    i++
  }
}

func dbInit() {
	db, err := gorm.Open(sqlite.Open"test.db", &gorm.Config{})
	if err != nil {
		panic("Cannot open DataBase(in func dbInit)")
	}

	db.AutoMigrate(&MLog{})
	defer db.Close()
}

func dbInsert(name string, text string) {
	db, err := gorm.Open("sqlite3", "test2.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInsert)")
	}

	db.Create(&MLog{Name: name, Text: text})
	defer db.Close()
}

func dbGet() []MLog {
	db, err := gorm.Open("sqlite3", "test2.sqlite3")
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
	//defer fmt.Println(myname, Isaid)

	dbInit()
	dbInsert(myname, Isaid)

	msg := dbGet()
  //[]msgを1つずつ表示
	display(msg)

}
