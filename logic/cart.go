package logic

import (
	"mozzarella-book/model"
)

func AddCart(cart model.Cart) (err error) {

	err = model.AddCart(cart)
	return
}

func DeleteCart(bookId, uid string) (err error) {
	err = model.DeleteCart(bookId, uid)
	return
}

func ShowCart(uid string) (carts []Cart, err error) {
	//查询用户购物车中有的书
	modelCarts, err := model.ShowCart(uid)
	if err != nil {
		return
	}

	//遍历购物车
	for _, cart := range modelCarts {
		//根据购物车的bookId找到具体收藏的书
		b, err := model.SearchBookWithBookId(cart.BookId)
		if err != nil {
			return
		}
		carts = append(carts, Cart{
			Wear:      cart.Wear,
			Transport: cart.Transport,
			Book:      b.ToLogicBook(),
		})
	}
	return
}
