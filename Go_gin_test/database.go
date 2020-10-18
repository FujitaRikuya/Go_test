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

//いい感じにText部分だけ表示させたい
func disp(msg []MLog) {
	i := 1

	for _, text := range msg {
		fmt.Println(i, "; Name = ", text.Name)
		fmt.Println("	Text = ", text.Text)
		i++
	}

}

func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInit)")
	}

	db.AutoMigrate(&MLog{})
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

/*
func dbGetOne(key_name string) MLog {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbGetOne")
	}
	defer db.Close()

	var result MLog
	db.First(&result, key_name)

	return result
}
*/

func main() {
	/*
		myname := "Rikuya"
		Isaid := "Hi there:)"
	*/

	dbInit()
	//dbInsert(myname, Isaid)

	fmt.Println("test.sqlite3")
	msg := dbGet()
	//fmt.Println(msg) 全部表示されてしまう

	//並べて表示してみる
	disp(msg)
}
