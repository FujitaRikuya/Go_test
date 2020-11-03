package main

import (
	"log"
	"net/http"

	"../Go_login_test/crypto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

//ユーザー名を管理するDBのモデル
type User struct {
	gorm.Model
	Uname string `form:"Uname" gorm:"unique;not null"`
	Pass  string `form:"Pass"`
}

func dbInit() {
	db, err := gorm.Open("sqlite3", "user.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInit)")
	}

	db.AutoMigrate(&User{})
	defer db.Close()
}

func dbInsert(uname string, pass string) {
	db, err := gorm.Open("sqlite3", "user.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbInsert)")
	}

	db.Create(&User{Uname: uname, Pass: pass})
	defer db.Close()
}

func dbGet() []User {
	db, err := gorm.Open("sqlite3", "user.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func dbGet")
	}

	var user []User
	db.Order("created_at desc").Find(&user)
	db.Close()

	return user
}

// ユーザー登録用、重複してたらエラーを返す？
func createUser(username string, password string) []error {
	EncryptedPass, _ := crypto.PasswordEncrypt(password)
	db, err := gorm.Open("sqlite3", "user.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func createUser)")
	}
	defer db.Close()

	// Insert処理
	if err := db.Create(&User{Uname: username, Pass: EncryptedPass}).GetErrors(); err != nil {
		return err
	}
	return nil
}

func getUser(uname string) User {
	db, err := gorm.Open("sqlite3", "user.sqlite3")
	if err != nil {
		panic("Cannot open DataBase(in func getUser)")
	}
	defer db.Close()

	var user User
	db.First(&user, "Uname = ?", uname)
	db.Close()

	return user
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("HTML/*.html")

	dbInit()

	router.GET("/", func(ctx *gin.Context) {
		user := dbGet()
		ctx.HTML(200, "loginPage.html", gin.H{
			"user": user,
		})
	})

	router.GET("/loginpage", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", gin.H{})
	})

	//dbに追加してリダイレクト
	router.POST("/signined", func(ctx *gin.Context) {
		uname := ctx.PostForm("Uname")
		pass := ctx.PostForm("Pass")

		if err := createUser(uname, pass); err != nil {
			ctx.HTML(http.StatusBadRequest, "Error.html", gin.H{"err": err})
		}
		//dbInsert(uname, pass)
		ctx.Redirect(302, "/")
	})

	router.POST("/login", func(ctx *gin.Context) {

		// DBから取得したユーザーパスワード(Hash)
		dbPassword := getUser(ctx.PostForm("Uname")).Pass
		log.Println(dbPassword)
		// フォームから取得したユーザーパスワード
		formPassword := ctx.PostForm("Pass")

		// ユーザーパスワードの比較
		if err := crypto.CompareHashAndPassword(dbPassword, formPassword); err != nil {
			log.Println("ログインできませんでした")
			ctx.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
			ctx.Abort()
		} else {
			log.Println("ログインできました")
			//c.Redirect(302, "/")

			//text := GetUlog() メッセージログを全部取り出す
			//ログイン後にチャットページに飛ぶ
			ctx.HTML(200, "Chat.html", gin.H{
				//"user": user, "text": text,
			})
		}
	})

	router.Run()
}
