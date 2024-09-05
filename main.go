package main

import (
	"fmt"
	"strconv"
	hanlder "test/Hanlder"

	"github.com/gin-gonic/gin"
)

func main() {
	var password = "sdasdasdasdawdwd"
	fmt.Print(password)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})
	r.GET("/add", Add)
	err := r.Run(":2001")
	if err != nil {
		panic(err)
	}
}

func Add(c *gin.Context) {
	num1 := c.Query("num1")
	num2 := c.Query("num2")
	a1, _ := strconv.Atoi(num1)
	a2, _ := strconv.Atoi(num2)
	sum := hanlder.AddFunc(a1, a2)
	c.JSON(200, gin.H{
		"result": sum,
	})
}
