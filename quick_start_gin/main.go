package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world This is gin server",
		})
	})

	engine.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"test": "OK"})
	})

	engine.Run(":3000")
}
