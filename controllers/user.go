package controllers

import (
    "github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/forms"
	"github.com/Rhodanthe1116/go-gin-boilerplate/auth"
    "github.com/Rhodanthe1116/go-gin-boilerplate/config"
    "github.com/Rhodanthe1116/go-gin-boilerplate/models"
	"net/http"
)

type UserController struct{}

func (h UserController) Signup(c *gin.Context) {
    var Payload forms.UserSignup
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // if item,_ := models.GetUserByPhone(Payload.Phone); item!=nil {
    //     c.JSON(http.StatusBadRequest, gin.H{"error": "Phone existed.", "user": item})
    //     return
    // }
    if err := HashPassword(&Payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    uuid := uuid.NewV4().String()
    Payload.UUID=uuid
    user := models.User{
        Phone: Payload.Phone,
        Password: Payload.Password,
        UUID: Payload.UUID,
    }
    if _,err := user.Signup(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, Payload)
}

func (h UserController) Login(c *gin.Context) {
    var Payload forms.UserLogin
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user,err := models.GetUserByPhone(Payload.Phone)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := CheckPassword(user.Password,Payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	config := config.GetConfig()
    jwtWrapper := auth.JwtWrapper{
		SecretKey:       config.GetString("jwt.secret"),
		Issuer:          "AuthService",
		ExpirationHours: config.GetInt64("jwt.expiration"),
	}
    signedToken, err := jwtWrapper.GenerateToken(Payload.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
    var token forms.UserToken
    token.Token = signedToken
    c.JSON(200,token)
}

