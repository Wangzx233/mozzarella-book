package model

func AddBook(book Book) (err error) {

	err = DB.Create(&book).Error
	return
}
