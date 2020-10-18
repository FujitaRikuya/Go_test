package dblib

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
func Disp(msg []MLog) {
	i := 1

	for _, text := range msg {
		fmt.Println(i, "; Name = ", text.Name)
		fmt.Println("	Text = ", text.Text)
		i++
	}

}

func DbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInit)")
	}

	db.AutoMigrate(&MLog{})
	defer db.Close()
}

func DbInsert(name string, text string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInsert)")
	}

	db.Create(&MLog{Name: name, Text: text})
	defer db.Close()
}

func DbGet() []MLog {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbGet")
	}
