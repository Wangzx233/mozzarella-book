package logic

import (
	"mozzarella-book/model"
	"strconv"
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
			ValueId: value.ValueId,
			Large:   image.Large,
			Medium:  image.Medium,
			Tiny:    image.Tiny,
		}

		modelValue.Images = append(modelValue.Images, modelImage)
	}

	return
}

// ToLogicBook 把model层book转换为logic层的book
func ToLogicBook(b model.Book) (LogicBook Book) {
	LogicBook = Book{
		BookId:    strconv.Itoa(int(b.ID)),
		Title:     b.Title,
		PublishBy: b.PublishBy,
		Edition:   b.Edition,
		Author:    b.Author,
	}

	for _, value := range b.Value {
		LogicBook.Value = append(LogicBook.Value, ToLogicValue(value))
	}

	return
}

func ToLogicValue(modelValue model.Value) (logicValue Value) {
	logicValue = Value{
		ValueId: modelValue.ID,
		BookId:  modelValue.BookId,
		Price:   modelValue.Price,
		Wear:    modelValue.Wear,
	}

	for _, image := range modelValue.Images {
		logicValue.Images = append(logicValue.Images, Images{
			Large:  image.Large,
			Medium: image.Medium,
			Tiny:   image.Tiny,
		})
	}
	return
}

func ToLogicCart(modelCart model.Cart) (logicCart Cart) {
	logicCart = Cart{
		Wear:      modelCart.Wear,
		Transport: modelCart.Transport,
	}
	return
}
