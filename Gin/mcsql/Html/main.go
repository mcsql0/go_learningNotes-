package main

import "github.com/gin-gonic/gin"

func main() {
	en := gin.Default()
	en.Static("/static","./static")
	en.LoadHTMLGlob("templates/*")

	en.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"title" : "Main website",
			"home" : "FPWcraft",
		})
	})


	en.Run(":80")
}
