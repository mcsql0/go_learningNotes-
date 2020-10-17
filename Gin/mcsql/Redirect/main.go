package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	en := gin.Default()
	//重定向
	en.GET("/ww", func(c *gin.Context) {
		c.Redirect(302,"http://www.baidu.com/")
	})
	//转发
	en.GET("/A", func(c *gin.Context) {
		c.Set("user", "xiaoming")
		c.Request.URL.Path = "/B"
		en.HandleContext(c)
	})

	en.GET("/B", func(c *gin.Context) {
		if v,ok := c.Get("user"); ok {
			c.String(http.StatusOK, v.(string))
		} else {
			c.String(http.StatusOK, "No")
		}
	})
	
	en.Run(":80")


}
