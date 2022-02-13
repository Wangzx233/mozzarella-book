package logic

import (
	"mozzarella-book/model"
)

func AddBook(book Book) (err error) {

	// 构建model层的book用于储存
	var moduleBook = model.Book{
		Title:     book.Title,
		PublishBy: book.PublishBy,
		Edition:   book.Edition,
		Value: model.Value{
			Price: book.Value.Price,
			Wear:  book.Value.Wear,
		},
	}

	// 因为model层的book的images要多一个bookId字段，所以需要单独赋值
	for _, image := range book.Images {
		moduleBook.Images = append(moduleBook.Images, model.Images{
			Large:  image.Large,
			Medium: image.Medium,
			Tiny:   image.Tiny,
		})

	}

	//添加
	err = model.AddBook(moduleBook)

	return
}
