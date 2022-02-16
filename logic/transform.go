package logic

import (
	"mozzarella-book/model"
)

// ToModelBook 把model层book转换为logic层的book
func (logicBook *Book) ToModelBook() (modelBook model.Book) {
	modelBook = model.Book{
		Title:     logicBook.Title,
		PublishBy: logicBook.PublishBy,
		Edition:   logicBook.Edition,
		Author:    logicBook.Author,
	}

	for _, value := range logicBook.Value {
		modelBook.Value = append(modelBook.Value, value.ToModelValue())
	}

	return
}

// ToModelValue logic层value转model层value
func (value *Value) ToModelValue() (modelValue model.Value) {
	modelValue = model.Value{
		BookId: value.BookId,
		Price:  value.Price,
		Wear:   value.Wear,
	}

	//images结构不同单独处理
	for _, image := range value.Images {
		modelImage := model.Images{
			BookId: value.BookId,
			Large:  image.Large,
			Medium: image.Medium,
			Tiny:   image.Tiny,
		}

		modelValue.Images = append(modelValue.Images, modelImage)
	}

	return
}
