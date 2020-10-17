package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	en := gin.Default()

	en.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ping")
	})

	en.Any("/help", func(c *gin.Context) {
		c.String(http.StatusOK, "Help")
	})

	en.NoRoute(func(c *gin.Context) {
		c.Request.URL.Path = "/ping"
		en.HandleContext(c)
	})

	gr := en.Group("/user")
	{
		gr.GET("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "Login")
		})
	}

	en.Run(":80")
}
