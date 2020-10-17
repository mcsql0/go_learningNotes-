package moder

import (
	"IDE/Gin/mcsql/Test/Book/utils"
)

type Book struct {
	BookID int `json:"book_id"`
	BookName string `json:"book_name"`
	BookTitle string `json:"book_title"`
	BookPrice string `json:"book_price"`
}

//添加图书
func AddBook (book_name string, book_title string, book_price string) error {

	//sql语句
	sql := "insert into books(book_name,book_title,book_price) values(?,?,?)"
	//预编译
	inStmt, err := utils.Db.Prepare(sql)
	if err != nil {
		panic(err)
		return err
	}
	//执行
	_, err = inStmt.Exec(book_name, book_title, book_price)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

//删除图书
func DelBook (book_id string) error {
	//sql语句
	sql := "delete from books where book_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sql)
	if err != nil {
		panic(err)
		return err
	}
	//执行
	_, err = inStmt.Exec(book_id)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

//查询图书
func QueryBook () ([]*Book,error) {
	sql := "select * from books"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		panic(err)
		return nil, err
	}
	var books []*Book
	for rows.Next() {
		var book_id int
		var book_name string
		var book_title string
		var book_price string
		rows.Scan(&book_id, &book_name, &book_title, &book_price)
		book := &Book{
			BookID: book_id,
			BookName: book_name,
			BookPrice: book_price,
			BookTitle: book_title,
		}
		books = append(books, book)
	}
	return books, err
}
