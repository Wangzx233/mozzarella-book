package route

import (
	"github.com/gin-gonic/gin"
	"mozzarella-book/api"
	"mozzarella-book/middleware"
)

func InitRoute() {
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"192.168.1.2"})
	if err != nil {
		return
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(200, "pong")
	})

	book := r.Group("/book", middleware.CheckUid)

	//添加书
	book.POST("/", api.AddBook)
	//添加磨损
	book.POST("/value", api.AddValue)
	//点击书
	book.GET("/click", api.ClickBook)
	//主界面和搜索
	book.GET("/search", api.GetBooks)
	//删除购物车中的书
	book.DELETE("/cart", api.DeleteCart)
	//添加书到购物车
	book.POST("/cart", api.AddCart)
	//查看购物车
	book.GET("/cart", api.ShowCart)

	r.Run(":8082")
}
