package logic

import (
	"mozzarella-book/model"
	"strconv"
)

// ModelBookToLogicBook 把model层book转换为logic层的book
func ModelBookToLogicBook(b model.Book) (LogicBook Book) {

	for _, image := range b.Images {
		LogicBook.Images = append(LogicBook.Images, Images{
			Large:  image.Large,
			Medium: image.Medium,
			Tiny:   image.Tiny,
		})
	}
	bookId := strconv.Itoa(int(b.ID))
	LogicBook.BookId = bookId
	LogicBook.Value = Value{
		Price: b.Value.Price,
		Wear:  b.Value.Wear,
	}
	LogicBook.Title = b.Title
	LogicBook.Edition = b.Edition
	LogicBook.PublishBy = b.PublishBy

	return
}