package model

import "gorm.io/gorm"

func AddValue(value Value) (err error) {
	err = DB.Create(&value).Error
	return
}

func UpdateValue(value Value) (err error) {
	err = DB.Model(Value{}).Where("book_id = ? and wear = ?", value.BookId, value.Wear).Updates(&value).Error
	DB.Model(&value).Association("Images").Replace(value.Images).Error()
	return
}

func CheckValue(bookId string, wear int) (b bool, value Value, err error) {
	err = DB.Preload("Images").Model(&Value{}).Where("book_id = ? and wear = ?", bookId, wear).First(&value).Error
	if err != nil {
		b = false
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	b = true
	return
}
