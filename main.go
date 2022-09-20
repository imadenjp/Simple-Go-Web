package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	id int `uri:"id" binding:"required"`
	name string `uri:"name" binding:"required,id"`
	msg string `uri:"msg" binding:"required"`
}

func main() {
	Gin := gin.Default()

	Gin.GET("/:id/:name", func(ctx *gin.Context){
		var person Person
		person.id = 100
		person.name = "moonliez"
		person.msg = "Hello World!"

		if err := ctx.ShouldBindUri(&person); err != nil {
			ctx.JSON(400, gin.H{"msg": err})
		}
		ctx.JSON(200, gin.H{
			"id": person.id,
			"name": person.name,
			"msg": person.msg,
		})
	})
	Gin.Run(":8000")
}