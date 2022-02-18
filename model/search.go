package model

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	limit = 20 //每页个数
)

func SearchBooksWithKeyword(keyword string, page int) (books []Book, err error) {
	err = DB.Preload("Value").Preload("Value.Images").Where("title LIKE ? OR publish_by LIKE ? OR "+
		"edition LIKE ? OR author LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Limit(limit).Offset((page - 1) * limit).Find(&books).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func SearchBooksWithKeywordAndWear(keyword, wear string) (books []Book, err error) {
	err = DB.Preload("Value").Preload("Value.Images").Where("(title LIKE ? OR publish_by LIKE ? OR "+
		"edition LIKE ? OR author LIKE ?) AND wear = ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", wear).Find(&books).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func SearchBookWithBookId(bookId string) (books Book, err error) {
	err = DB.Preload("Value").Preload("Value.Images").Model(&Book{}).Where("id=?", bookId).First(&books).Error
	if err == gorm.ErrRecordNotFound {
		return Book{}, nil
	}
	return
}

func GetAllBooks(page, seed int, uid string) (books []Book, err error) {
	var otherBook []Book
	var userClickRecordCount int64
	//找出6个用户点击过的
	err = DB.Preload("Value").Preload("Value.Images").Where("id IN (select book_id from user_click_records where uid = ?)", uid).Order(fmt.Sprintf("RAND(%d)", seed)).Limit(6).Offset(6 * (page - 1)).Find(&books).Count(&userClickRecordCount).Error
	//找出limit-6个用户没点击过的
	err = DB.Preload("Value").Preload("Value.Images").Where("id NOT IN (select book_id from user_click_records where uid = ?)", uid).Order(fmt.Sprintf("RAND(%d)", seed)).Limit(int(limit - userClickRecordCount)).Offset((page - 1) * int(limit-userClickRecordCount)).Find(&otherBook).Error
	books = append(books, otherBook...)

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}
