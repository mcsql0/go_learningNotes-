package moder

import (
	"fmt"
	"testing"
)

func TestBook_AddBook(t *testing.T) {
	if err:= AddBook("钢蛋是咋样练成的", "book", "50"); err != nil {
		panic(err)
	}
}

func TestQueryBook(t *testing.T) {
	if books, err := QueryBook(); err != nil {
		panic(err)
	} else {
		for _,v := range books{
			fmt.Println(v.BookTitle)
		}
	}
}