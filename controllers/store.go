package controllers

import (
    "github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/forms"
	"github.com/Rhodanthe1116/go-gin-boilerplate/auth"
    "github.com/Rhodanthe1116/go-gin-boilerplate/config"
	"net/http"
)

type StoreController struct{}

func (h StoreController) Signup(c *gin.Context) {
    var Payload forms.StoreSignup
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // TODO: if(db.get(Payload.Phone)) then duplicated;
    if err := HashPassword(&Payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    uuid := uuid.NewV4().String()
    Payload.UUID=uuid
    // TODO: db.insert(Payload.UUID,Payload.Name,Payload.Password,Payload.Phone,Payload.Address);
	// c.String(http.StatusOK, "Success")
    c.JSON(200, Payload)
}

func (h StoreController) Login(c *gin.Context) {
    var Payload forms.StoreLogin
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // TODO: store=db.get(Payload.Phone);
    if err := CheckPassword("$2a$14$hNbyJM1JPwCTnd4Yx3AGG.lITDqeBOrEhzh2/fs3zA2lJ7rTwn12G",Payload.Password); err != nil {
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
    var token forms.StoreToken
    token.Token = signedToken
    c.JSON(200,token)
}

func (h StoreController) Profile(c *gin.Context){
    phone,_ := c.Get("phone")
    // TODO: store=db.get(phone);
    var Profile forms.StoreProfile
    Profile.Name="giver"
    Profile.Phone=phone.(string)
    Profile.Address="abcdefg"
    Profile.QrCode="https://example.com/"
    c.JSON(200,Profile)
}

