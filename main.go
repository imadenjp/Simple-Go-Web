package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func pingPong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", helloWorld)
	router.GET("/ping", pingPong)

	router.Run("0.0.0.0:8000")
}