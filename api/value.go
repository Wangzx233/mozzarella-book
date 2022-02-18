package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-book/logic"
)

func AddValue(c *gin.Context) {
	var value logic.Value
	err := c.BindJSON(&value)
	if err != nil {
		log.Println("json bind err : ", err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "parameter error",
		})
		return
	}

	err = logic.AddValue(value)
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
