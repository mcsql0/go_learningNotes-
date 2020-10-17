package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	// Default 已经连接了 Logger 和 Recovery 中间件
	en := gin.Default()
	
	en.GET("/ping", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"user" : "颤三",
		})
	})

	en.GET("user/:name", func(context *gin.Context) {
		name := context.Param("name")
		//http.StatusOK = 200
		//将给定的字符串写入响应体。
		context.String(http.StatusOK,"Hello %s",name)
	})

	//* 通用匹配
	en.GET("user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + "is" + action
		context.String(200,message)
	})

	//查询字符串参数
	//127.0.0.1:8080/help?firstname=Jane&lastname=Doe
	//相应 Hello Jane Doe
	//127.0.0.1:8080/help
	//相应 Hello 默认值
	en.GET("/help", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname","默认值")
		name := context.Query("lastname")
		context.String(200, "Hello %s %s", firstname, name)
	})

	//Multipart/Urlencoded 表单
	en.POST("Post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "默认值")

		context.JSON(200, gin.H{
			"status" : "posted",		//他有可能排到后面
			"message" : message,
			"nick" : nick,
		})
	})

	//GET+POST
	en.POST("/post", func(context *gin.Context) {
		id := context.Query("id")		//获取GET字段值
		page := context.DefaultQuery("page","0")	//获取URL字段，如果没有默认使用后面的默认值
		name := context.PostForm("name")
		message := context.PostForm("message")		//表单字段查询

		context.JSON(200,gin.H{
			id : id,
			page : page,
			name : name,
			message : message,
		})
	})

	//文件上传
	en.POST("/upload", func(context *gin.Context) {
		//单文件
		file, err := context.FormFile("file")
		if err != nil {
			log.Panicln(err)
			return
		}
		log.Println(file.Filename)
		context.String(200,"File Name: %s", string(file.Filename))
	})

	//多文件
	en.POST("/uploads", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			log.Println(err)
			return
		}
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println("upload ---> file Name : %s", file.Filename)
		}
		context.String(200,"upload file num = %s", len(files))
	})

	//组路由
	v1 := en.Group("/v1")
	{
		v1.POST("/login", func(context *gin.Context) {
			context.String(200,"login")
		})

		v1.POST("/register", func(context *gin.Context) {
			context.String(200, "register")
		})

		v1.GET("/help", func(context *gin.Context) {
			context.String(200,"help")
		})

		v1.DELETE("/del", func(context *gin.Context) {
			context.String(200,"del")
		})
	}



	_ = en.Run()
}
