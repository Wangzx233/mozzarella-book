package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-book/logic"
)

// GetBooks 通过keyword搜索主页商品，keyword为空即所有商品
func GetBooks(c *gin.Context) {
	keyword := c.Query("keyword")

	books, err := logic.GetBooksWithKeyword(keyword)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "parameter err",
			"data":   books,
		})
	} else {
		c.JSON(200, gin.H{
			"status": 10000,
			"info":   "success",
			"data":   books,
		})
	}
}
