package logic

import (
	"log"
	"mozzarella-book/model"
)

func AddBook(book Book) (err error) {

	//添加
	err = model.AddBook(book.ToModelBook())

	return
}

func ClickBook(bookId, uid string) bool {
	err := model.ClickBook(bookId, uid)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
