package model

import "gorm.io/gorm"

func AddCart(cart Cart) error {
	err := DB.Where("book_id = ? AND uid = ?", cart.BookId, cart.Uid).First(&cart).Error

	if err == gorm.ErrRecordNotFound {
		err = DB.Create(&cart).Error
	}

	return err
}

func DeleteCart(bookId, uid string) error {
	err := DB.Where("book_id = ? AND uid = ?", bookId, uid).Delete(&Cart{}).Error
	return err
}

func ShowCart(uid string) (carts []Cart, err error) {
	err = DB.Model(&Cart{}).Where("uid=?", uid).Find(&carts).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}
