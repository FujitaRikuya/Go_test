package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	zikoku := time.Now()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"zikoku": zikoku})
	})

	router.Run()
}
