package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

var en = gin.Default()

type user struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
//中间件：判断用户是否登录
func isLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* 内存级 */
		//if _,ok := c.Get("user"); !ok {
		//
		//}
		//if _,ok := c.Get("user"); !ok {
		//	//c.Redirect(http.StatusMovedPermanently,"/Login")
		//	c.Request.URL.Path = "/Login"
		//	en.HandleContext(c)
		//} else {
		//	c.Next()		//执行后续操作
		//}
		/* 用Cookie 判断是否登录 （不安全） */
		if _, err := c.Cookie("wwName"); err != nil {
			c.Request.URL.Path = "/Login"
			en.HandleContext(c)
		} else {
			c.Next()		//执行后续操作
		}

	}
}

func homeHandler(c *gin.Context) {
	if k,b := c.Get("user"); b {
		c.String(http.StatusOK,"User:%s",k)
	}
}

func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", gin.H{

	})
}

func loginPostHandler(c *gin.Context) {
	var u user
	if err := c.ShouldBindWith(&u,binding.Form); err == nil {
		c.Set("user",u)
		/* 设置Cookie */
		c.SetCookie("wwName","郭阳阳",9000,"127.0.0.1","",true,true)
		c.Request.URL.Path = "/home"
		en.HandleContext(c)
		//c.Redirect(http.StatusMovedPermanently,"/home")
		fmt.Println("登录成功")
	} else {
		fmt.Println(err)
	}

	fmt.Println("ssss")

}

func main()  {

	en.LoadHTMLGlob("templates/**/*")

	en.GET("/home", isLogin(), homeHandler)
	en.GET("/Login",loginHandler)
	en.POST("/Login",loginPostHandler)

	en.Run(":80")
}
