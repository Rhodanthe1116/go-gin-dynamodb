package controllers

import (
    "strconv"

    "github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/forms"
	"github.com/Rhodanthe1116/go-gin-boilerplate/auth"
    "github.com/Rhodanthe1116/go-gin-boilerplate/config"
    "github.com/Rhodanthe1116/go-gin-boilerplate/models"
	"net/http"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

type StoreController struct{}

func (h StoreController) Signup(c *gin.Context) {
    var Payload forms.StoreSignup
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if _, err := strconv.ParseInt(Payload.Phone,10,64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Phone is not a valid phone number."})
        return
    }
    if item,_ := models.GetStoreByPhone(Payload.Phone); item!=nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Phone existed.", "store": item})
        return
    }
    if err := HashPassword(&Payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    uuid := uuid.NewV4().String()
    Payload.UUID=uuid
    store := models.Store{
        Phone: Payload.Phone,
        Password: Payload.Password,
        UUID: Payload.UUID,
        Name: Payload.Name,
        Address: Payload.Address,
    }
    if _,err := store.Signup(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, Payload)
}

func (h StoreController) Login(c *gin.Context) {
    var Payload forms.StoreLogin
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    store,err := models.GetStoreByPhone(Payload.Phone)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := CheckPassword(store.Password,Payload.Password); err != nil {
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
    store,err := models.GetStoreByPhone(phone.(string))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	config := config.GetConfig()
    key := config.GetString("hmac.key")
    mac := hmac.New(sha256.New,[]byte(key))
    mac.Write([]byte(store.Phone))
    MAC := hex.EncodeToString(mac.Sum(nil))
    var Profile forms.StoreProfile
    Profile.Name=store.Name
    Profile.Phone=store.Phone
    Profile.Address=store.Address
    Profile.UUID=store.UUID
    Profile.QrCode=store.Phone+"||"+MAC
    c.JSON(200,Profile)
}

