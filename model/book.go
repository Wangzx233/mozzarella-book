package model

import (
	"errors"
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

// CheckWear 查看书是否已存在,true为存在
func CheckWear(bookId string, wear int) (exited bool, err error) {
	exited = true
	err = DB.Model(&Value{}).Where("book_id = ? and wear = ?", bookId, wear).First(&Value{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			exited = false
			err = nil
		}
		return
	}
	return
}

func ClickBook(bookId, uid string) (err error) {
	err = DB.Model(&Book{}).Where("id = ?", bookId).First(&Book{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			var UCR = UserClickRecord{
				Uid:    uid,
				BookId: bookId,
			}
			err = DB.Create(&UCR).Error
		}
	} else {
		err = errors.New("record exited")
	}

	return
}
