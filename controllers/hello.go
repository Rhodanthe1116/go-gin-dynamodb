package controllers

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func (h HelloController) Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
