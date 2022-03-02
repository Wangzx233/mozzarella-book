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
	err = DB.Where("uid=?", uid).Find(&carts).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func CheckCart(bookId, uid string) (exited bool, err error) {
	exited = true
	err = DB.Model(&Cart{}).Where("book_id = ? and uid = ?", bookId, uid).First(&Cart{}).Error
	if err == gorm.ErrRecordNotFound {
		exited = false
		err = nil
	}

	return
}

func (cart *Cart) UpdateCart() (err error) {
	err = DB.Model(Cart{}).Where("book_id = ?", cart.BookId).Updates(&cart).Error
	return
}
