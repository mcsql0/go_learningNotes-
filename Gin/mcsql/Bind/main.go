package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Login struct {
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	en := gin.Default()

	en.POST("/ShouldBind", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBind(&json); err == nil  {
			if json.User == "manu" && json.Password == "123" {
				context.String(200,"登录成功")
			} else {
				context.JSON(200, gin.H{
					"登录错误！" : json.User,
				})
			}
		} else {
			context.String(200, "err = %s",err)
		}
	})

	//{ "user": "manu", "password": "123" }
	en.POST("/ShouldBindWith", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindWith(&json,binding.JSON); err == nil  {
			if json.User == "manu" && json.Password == "123" {
				context.String(200,"登录成功")
			} else {
				context.JSON(200, gin.H{
					"登录错误！" : json.User,
				})
			}
		} else {
			context.String(200, "err = %s",err)
		}
	})

	//只绑定 url 查询字符串
	//ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据。
	en.GET("/ShouldBindQuery", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindQuery(&json); err != nil {
			if json.User == "manu" && json.Password == "123" {
				context.String(200,"登录成功")
			} else {
				context.JSON(200, gin.H{
					"登录错误！" : json.User,
				})
			}
		} else {
			context.String(200, "err = %s",err)
		}
	})

	en.Run(":80")
}
