package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//var text string

	zikoku := time.Now()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	/*
		router.POST("/pushed", func(ctx *gin.Context) {
			text = ctx.PostForm("Chattext")
			file, err := os.Open(`C:\Users\ayukir\go\Go_gin_Chat\test.txt`)

			if err != nil {
				log.Fatal(err)
			}

			file.Write(([]byte)(text))
			defer file.Close()

			ctx.Redirect(302, "/")
		})
	*/

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "Chat.html", gin.H{
			"zikoku": zikoku, //"text": text,
		})
	})

	router.Run()
}
