package main

import (
	"github.com/gin-gonic/gin"
	"github.com/songjiayang/gin-mod/web/model"
	"github.com/songjiayang/go-mod/lib/math"
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

	r.GET("/total", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"i": 1,
			"j": 2,
			"total":math.Sum(1,2),
		})
	})
	r.Run()
}
