package logic

import (
	"log"
	"mozzarella-book/model"
)

func GetBooksWithKeywordAndWear(keyword, wear string) (books []Book, err error) {
	var modelBooks []model.Book
	if keyword == "" && wear == "" {
		modelBooks, err = model.GetAllBooks()
	} else {
		if wear == "" {
			modelBooks, err = model.SearchBooksWithKeyword(keyword)
		} else {
			modelBooks, err = model.SearchBooksWithKeywordAndWear(keyword, wear)
		}
	}

	log.Println(modelBooks)
	for _, book := range modelBooks {
		//把model层book转换为logic层的book并加入
		books = append(books, book.ToLogicBook())
	}
	return
}
