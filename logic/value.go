package logic

import (
	"errors"
	"log"
	"mozzarella-book/model"
)

func AddValue(value Value) (err error) {
	var rightWears = [4]int{100, 90, 70, 50}
	var r = false
	for _, rightWear := range rightWears {
		if value.Wear == rightWear {
			r = true
			break
		}
	}

	if !r {
		return errors.New("wear not allow")
	}

	//检查书是否存在
	b, err := model.CheckBook(value.BookId)
	if err != nil {
		log.Println(err)
		return
	}
	if !b {
		err = errors.New("book_id not existed")
		return
	}

	//检查书是否已有相同磨损存在
	//todo:可优化
	b, modelValue, err := model.CheckValue(value.BookId, value.Wear)
	if err != nil {
		log.Println(err)
		return
	}
	value.ValueId = modelValue.ID
	if b {
		err = model.UpdateValue(value.ToModelValue())
		return
	}

	//添加
	err = model.AddValue(value.ToModelValue())
	return
}
