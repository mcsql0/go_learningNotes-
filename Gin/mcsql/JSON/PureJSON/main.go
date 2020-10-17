package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	en := gin.Default()

	// 提供 unicode 实体 将特殊字符转为unicode
	/* 输出: {"html":"\u003cb\u003eHello, world!\u003c/b\u003e"} */
	en.GET("/json", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"html" : "<b>Hello, world!</b>",
		})
	})

	//提供字面字符 不改变特殊字符
	/* 输出: {"html":"<b>Hello, world!</b>"} */
	en.GET("/purejson", func(context *gin.Context) {
		context.PureJSON(200, gin.H{
			"html" : "<b>Hello, world!</b>",
		})
	})

	en.Run(":80")
}
