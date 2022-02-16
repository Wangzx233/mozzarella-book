package model

import "gorm.io/gorm"

func AddValue(value Value) (err error) {
	err = DB.Create(&value).Error
	return
}

func CheckValue(bookId string, wear int) (bool, error) {
	err := DB.Model(&Value{}).Where("book_id = ? and wear = ?", bookId, wear).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
