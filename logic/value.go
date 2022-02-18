package logic

import (
	"errors"
	"log"
	"mozzarella-book/model"
)

func AddValue(value Value) (err error) {
	//检查书是否存在
	b, err := model.CheckBook(value.BookId)
	if err != nil {
		log.Println(err)
		return
	}
	if !b {
		err = errors.New("book_id existed")
		return
	}

	//检查书是否已有相同磨损存在
	b, err = model.CheckValue(value.BookId, value.Wear)
	if err != nil {
		log.Println(err)
		return
	}
	if b {
		err = errors.New("wear existed")
		return
	}

	//添加
	err = model.AddValue(value.ToModelValue())
	return
}
