package logic

import "mozzarella-book/model"

func GetBooksWithKeyword(keyword string) (books []Book, err error) {
	var modelBooks []model.Book
	if keyword == "" {
		modelBooks, err = model.GetAllBooks()
	} else {
		modelBooks, err = model.SearchBooksWithKeyword(keyword)
	}

	for _, book := range modelBooks {
		//把model层book转换为logic层的book
		logicBook := ModelBookToLogicBook(book)

		books = append(books, logicBook)
	}
	return
}
