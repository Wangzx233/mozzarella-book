package model

import (
	"gorm.io/gorm"
)

func SearchBooksWithKeyword(keyword string) (books []Book, err error) {
	err = DB.Preload("Images").Where("title LIKE ? OR publish_by LIKE ? OR "+
		"edition LIKE ? OR author LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Find(&books).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func SearchBooksWithKeywordAndWear(keyword, wear string) (books []Book, err error) {
	err = DB.Preload("Images").Where("(title LIKE ? OR publish_by LIKE ? OR "+
		"edition LIKE ? OR author LIKE ?) AND wear = ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", wear).Find(&books).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func SearchBookWithBookId(bookId string) (books Book, err error) {
	err = DB.Preload("Images").Model(&Book{}).Where("id=?", bookId).First(&books).Error
	if err == gorm.ErrRecordNotFound {
		return Book{}, nil
	}
	return
}

func GetAllBooks() (books []Book, err error) {
	err = DB.Preload("Images").Find(&books).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}
