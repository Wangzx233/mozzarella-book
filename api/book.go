package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-book/logic"
)

func AddBook(c *gin.Context) {
	var book logic.Book
	err := c.BindJSON(&book)
	if err != nil {
		log.Println("json bind err : ", err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "internal server error",
		})
		return
	}

	err = logic.AddBook(book)
	if err != nil {
		log.Println(err)
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
