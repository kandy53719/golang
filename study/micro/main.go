package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var r = gin.Default()

	r.LoadHTMLFiles("templates/index.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"测试": "ok",
		})
	})

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"测试": "ok",
		})
	})

	r.Run(":8080")
}
