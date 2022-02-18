package middleware

import "github.com/gin-gonic/gin"

func CheckUid(c *gin.Context) {
	uid := c.GetHeader("uid")
	if uid == "" {
		c.JSON(401, gin.H{
			"status": 50000,
			"info":   "无授权",
		})
		c.Abort()
	} else {
		c.Set("uid", uid)
		c.Next()
	}
}
