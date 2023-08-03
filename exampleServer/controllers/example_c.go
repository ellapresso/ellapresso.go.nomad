package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	fmt.Println("pong")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
