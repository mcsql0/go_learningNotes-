package main

import (
	"IDE/Gin/mcsql/Test/Book/moder"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

var en = gin.Default()

func main()  {

	en.LoadHTMLGlob("./mcsql/Test/Book/templates/*")
	
	en.GET("/", indexHandler)
	en.GET("/book/new",newBookHandler)
	en.POST("/new_book",addBookHandler)
	en.GET("/book/delete",delBookHandler)

	en.Run(":80")
}

func indexHandler(c *gin.Context) {
	books, err := moder.QueryBook()
	if err != nil {
		panic(books)
		return
	}

	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"data" : books,
	})
}

func newBookHandler(c *gin.Context)  {
	c.HTML(http.StatusOK,"new_book.html",nil)
}

func addBookHandler(c *gin.Context)  {
	var b moder.Book
	if err := c.ShouldBindWith(&b, binding.Form); err == nil {
		//数据库操作
		if err := moder.AddBook(b.BookName, b.BookTitle, b.BookPrice); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err" : err,
			})
		} else {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	} else {
		panic(err)
	}
}

func delBookHandler(c *gin.Context)  {
	id := c.Query("id")
	if err := moder.DelBook(id); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"err" : err,
		})
	} else {
		c.Request.URL.Path = "/"
		en.HandleContext(c)
	}
}
