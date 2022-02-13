package model

import (
	"gorm.io/gorm"
)

func SearchBooksWithKeyword(keyword string) (books []Book, err error) {
	err = DB.Where("title LIKE ? OR publish_by LIKE ? OR "+
		"edition LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Find(&books).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func SearchBookWithBookId(bookId string) (books Book, err error) {
	err = DB.Model(&Book{}).Where("book_id=?", bookId).First(&books).Error
	if err == gorm.ErrRecordNotFound {
		return Book{}, nil
	}
	return
}

func GetAllBooks() (books []Book, err error) {
	err = DB.Find(&books).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}
