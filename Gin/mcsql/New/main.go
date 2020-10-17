package main

import "github.com/gin-gonic/gin"

func main()  {
	//创建一个默认的没有任何中间件的路由
	en := gin.New()

	// 全局中间件
	// Logger 中间件将写日志到 gin.DefaultWriter ,即使你设置 GIN_MODE=release 。
	// 默认 gin.DefaultWriter = os.Stdout
	en.Use(gin.Logger())

	//Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	en.Use(gin.Recovery())

	//每个路由的中间件, 你能添加任意数量的中间件
	//en.GET("benchmark")

	//授权组
	authorized := en.Group("/")
	authorized.Use(gin.Logger())
	{
		authorized.POST("/login", func(context *gin.Context) {
			context.String(200,"login")
		})
	}

	en.Run(":80")
}
