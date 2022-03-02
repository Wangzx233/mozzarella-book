package logic

import (
	"errors"
	"log"
	"mozzarella-book/model"
)

func AddCart(cart model.Cart) (err error) {
	// 检测wear是否存在
	wearExited, err := model.CheckWear(cart.BookId, cart.Wear)
	if err != nil {
		log.Println("check wear err : ", err)
		return errors.New("err")
	}
	// 检查book是否存在
	bl, err := model.CheckBook(cart.BookId)
	// book和对应磨损都存在才能操作
	if bl && wearExited {
		var exited = true
		// 检查是否已经收藏相同都书
		exited, err = model.CheckCart(cart.BookId, cart.Uid)
		if err != nil {
			log.Println(err)
			return
		}
		// 如果是相同都书就更新否则就新建
		if exited {
			//todo:更新cart
			err = cart.UpdateCart()
		} else {
			err = model.AddCart(cart)
		}
	} else {
		return errors.New("book or wear don't exited")
	}
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
	log.Println(uid)
	//遍历购物车
	for _, cart := range modelCarts {
		//根据购物车的bookId找到具体收藏的书
		b, err := model.SearchBookWithBookId(cart.BookId)
		if err != nil {
			break
		}
		carts = append(carts, Cart{
			Wear:      cart.Wear,
			Transport: cart.Transport,
			Book:      ToLogicBook(b),
		})
	}
	return
}
