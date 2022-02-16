package model

import (
	"mozzarella-book/logic"
	"strconv"
)

// ToLogicBook 把model层book转换为logic层的book
func (b *Book) ToLogicBook() (LogicBook logic.Book) {
	LogicBook = logic.Book{
		BookId:    strconv.Itoa(int(b.ID)),
		Title:     b.Title,
		PublishBy: b.PublishBy,
		Edition:   b.Edition,
		Author:    b.Author,
	}

	for _, value := range b.Value {
		LogicBook.Value = append(LogicBook.Value, value.ToLogicValue())
	}

	return
}

func (modelValue *Value) ToLogicValue() (logicValue logic.Value) {
	logicValue = logic.Value{
		BookId: modelValue.BookId,
		Price:  modelValue.Price,
		Wear:   modelValue.Wear,
	}

	for _, image := range modelValue.Images {
		logicValue.Images = append(logicValue.Images, logic.Images{
			Large:  image.Large,
			Medium: image.Medium,
			Tiny:   image.Tiny,
		})
	}
	return
}

func (modelCart *Cart) ToLogicCart() (logicCart logic.Cart) {
	logicCart = logic.Cart{
		Uid:       modelCart.Uid,
		BookId:    modelCart.BookId,
		Wear:      modelCart.Wear,
		Transport: modelCart.Transport,
	}
	return
}
