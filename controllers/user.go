package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/forms"
	"github.com/vsouza/go-gin-boilerplate/models"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func (u UserController) Signup(c *gin.Context) {
	// Example for binding JSON ({"user": "manu", "password": "123"})
	var userPayload forms.UserSignup
	if err := c.ShouldBindJSON(&userPayload); err != nil {
		fmt.Print(userPayload)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "test log", "user": userPayload})
	fmt.Print(userPayload)

	user, err := userModel.Signup(userPayload)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
	c.Abort()
	return
}
