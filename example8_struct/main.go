package main

import (
	"fmt"
	"hellogo/example8_struct/model"
)

func main() {
	var golang model.Book
	golang.Title = "Go"
	golang.Author = "CSI Professional"
	golang.Subtitle = "tutarial"
	golang.Price = 199.99
	golang.Book.BookCode = 001
	golang.Book.BookName = "CSI BOOK"
	golang.Book.BookCreated = "26/09/2563"

	fmt.Println(golang)

	golang2 := model.Book{
		Title:    "Go",
		Author:   "CSI Professional",
		Subtitle: "tutarial",
		Price:    199.99,
		Book: model.BookList{
			BookCode:    001,
			BookName:    "CSI BOOK",
			BookCreated: "26/09/2563",
		},
	}
	fmt.Println(golang2)
}
