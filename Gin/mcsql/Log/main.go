package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main()  {
	en := gin.Default()

	file,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)

	// 如果你需要同时写入日志文件和控制台上显示，使用下面代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	en.GET("/ping", func(context *gin.Context) {
		context.String(200,"ping")
	})

	en.Run(":80")
}
