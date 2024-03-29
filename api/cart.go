package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-book/logic"
	"mozzarella-book/model"
)

func DeleteCart(c *gin.Context) {
	var cart model.Cart
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		log.Println("json bind err : ", err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "json error",
		})
		return
	}

	uid, _ := c.Get("uid")
	cart.Uid = uid.(string)
	err = logic.DeleteCart(cart.BookId, cart.Uid)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "internal server error",
		})
	} else {
		c.JSON(200, gin.H{
			"status": 10000,
			"info":   "success",
		})
	}
}

func AddCart(c *gin.Context) {
	var cart model.Cart
	err := c.BindJSON(&cart)
	if err != nil {
		log.Println("json bind err : ", err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "json error",
		})
		return
	}

	//获取uid
	uid, _ := c.Get("uid")
	cart.Uid = uid.(string)

	err = logic.AddCart(cart)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   fmt.Sprint(err),
		})
	} else {
		c.JSON(200, gin.H{
			"status": 10000,
			"info":   "success",
		})
	}
}

func ShowCart(c *gin.Context) {
	//todo: 获取uid
	uid, _ := c.Get("uid")

	carts, err := logic.ShowCart(uid.(string))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "internal server error",
			"data":   carts,
		})
	} else {
		c.JSON(200, gin.H{
			"status": 10000,
			"info":   "success",
			"data":   carts,
		})
	}

}
