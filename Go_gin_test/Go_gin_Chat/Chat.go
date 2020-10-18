package main

import (
	"time"

	"github.com/gin-gonic/gin"
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
	db, err := gorm.Open("sqlite3", "msglog.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInit)")
	}

	db.AutoMigrate(&MLog{})
	defer db.Close()
}

func dbInsert(name string, text string) {
	db, err := gorm.Open("sqlite3", "msglog.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInsert)")
	}

	db.Create(&MLog{Name: name, Text: text})
	defer db.Close()
}

func dbGet() []MLog {
	db, err := gorm.Open("sqlite3", "msglog.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbGet")
	}

	var mlog []MLog
	db.Order("created_at desc").Find(&mlog)
	db.Close()

	return mlog
}

func main() {
	zikoku := time.Now()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	router.GET("/", func(ctx *gin.Context) {
		text := dbGet()
		ctx.HTML(200, "Chat.html", gin.H{
			"zikoku": zikoku, "text": text,
		})
	})

	//dbに追加してリダイレクト
	router.POST("/pushed", func(ctx *gin.Context) {
		name := ctx.PostForm("Name")
		text := ctx.PostForm("ChatText")
		dbInsert(name, text)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
