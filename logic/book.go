package logic

import (
	"mozzarella-book/model"
)

func AddBook(book Book) (err error) {

	//添加
	err = model.AddBook(book.ToModelBook())

	return
}
