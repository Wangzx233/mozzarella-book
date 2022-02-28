package logic

import (
	"math/rand"
	"mozzarella-book/model"
	"strconv"
)

func GetBooksWithKeyword(keyword, page, seed, uid string) (books []Book, sd int, err error) {
	var modelBooks []model.Book
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return
	}

	if keyword == "" {
		//如果是第一页或者没有seed，就生成seed
		if pageNumber == 1 || seed == "" {
			sd = rand.Intn(100)
		} else {
			sd, _ = strconv.Atoi(seed)
		}
		modelBooks, err = model.GetAllBooks(pageNumber, sd, uid)
	} else {
		modelBooks, err = model.SearchBooksWithKeyword(keyword, pageNumber)
	}

	for _, book := range modelBooks {

		//把model层book转换为logic层的book并加入
		books = append(books, ToLogicBook(book))
	}
	return
}
