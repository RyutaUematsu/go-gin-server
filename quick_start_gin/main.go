package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//UserAgentを取得するミドルウェアの使用
	ua := ""
	server.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})
	server.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message":    "hello world This is gin server",
			"User Agent": ua,
		})
	})

	server.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"test": "OK"})
	})

	server.Run(":3000")
}
