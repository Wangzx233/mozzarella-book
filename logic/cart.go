package logic

import (
	"mozzarella-book/model"
)

func AddCart(bookId, uid string) (err error) {
	//if num == "" {
	//	err = model.AddCart(bookId, uid)
	//} else {
	//	n, err := strconv.Atoi(num)
	//	if err != nil {
	//		return err
	//	}
	//
	//	err = model.AddCart(bookId, uid, uint(n))
	//}
	err = model.AddCart(bookId, uid)
	return
}

func DeleteCart(bookId, uid string) (err error) {
	err = model.DeleteCart(bookId, uid)
	return
}

func ShowCart(uid string) (books []Book, err error) {
	//查询用户购物车中有的书
	carts, err := model.ShowCart(uid)
	if err != nil {
		return nil, err
	}

	//遍历购物车
	for _, cart := range carts {
		//根据购物车的bookId找到具体收藏的书
		b, err := model.SearchBookWithBookId(cart.BookId)
		if err != nil {
			return nil, err
		}

		//把model层book转换为logic层的book
		tmpBook := ModelBookToLogicBook(b)

		//加入到结果
		books = append(books, tmpBook)
	}
	return
}
