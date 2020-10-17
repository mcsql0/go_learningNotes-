package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	en := gin.Default()

	//你也可以使用自己的 SecureJSON 前缀
	//en.SecureJsonPrefix(")]}',\n")
	
	en.GET("/SecureJSON", func(context *gin.Context) {
		names := []string{"1","2","3","4","5"}
		//将输出：while(1);["lena","austin","foo"]
		context.SecureJSON(http.StatusOK,names)
	})

	en.Run(":80")
}
