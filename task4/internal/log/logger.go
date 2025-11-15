package mylog

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL) // 打印出c.Request.URL.
		c.Next()
		fmt.Println(c.Writer.Status())

	}
}
