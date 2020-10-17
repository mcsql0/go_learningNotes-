package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context)  {
	fmt.Println("中间件执行前！")
	c.Next()
	fmt.Println("中间件执行后！")
}

func main()  {
	c := gin.Default()

	c.Use(LoginHandler)
	{
		c.GET("/login", func(co *gin.Context) {
			fmt.Println("/login 路由执行")
			co.String(200,"login")
		})

		c.GET("/res", func(co *gin.Context) {
			fmt.Println("/res 路由执行")
			co.String(200,"res")
		})
	}

	c.Run(":80")
}
