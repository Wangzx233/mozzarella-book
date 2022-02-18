package model

import (
	"gorm.io/gorm"
)

func AddBook(book Book) (err error) {

	err = DB.Create(&book).Error
	return
}

// CheckBook 查看书是否已存在,true为存在
func CheckBook(bookId string) (bool, error) {

	err := DB.Model(&Book{}).Where("id = ?", bookId).First(&Book{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func ClickBook(bookId, uid string) (err error) {
	var UCR = UserClickRecord{
		Uid:    uid,
		BookId: bookId,
	}
	err = DB.Create(&UCR).Error

	return
}