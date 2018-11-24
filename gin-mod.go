package main

import (
	"github.com/gin-gonic/gin"
	"github.com/songjiayang/gin-mod/model"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, model.NewUser("josh"))
	})
	r.Run()
}
