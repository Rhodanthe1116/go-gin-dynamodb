package controllers

import (
    "time"
    "strings"

	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/forms"
	"github.com/Rhodanthe1116/go-gin-boilerplate/models"
    "github.com/Rhodanthe1116/go-gin-boilerplate/config"
	"net/http"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

type RecordController struct{}

func (h RecordController) Record(c *gin.Context) {
    var Payload forms.RecordStoreId
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    userPhone,_ := c.Get("phone")

    storeIdMac := Payload.StoreId
    strs := strings.Split(storeIdMac, "||")

    if len(strs)<2 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Store id format error"})
        return
    }

    MAC := strs[len(strs)-1]
    storePhone := strings.Join(strs[:len(strs)-1],"||")

	config := config.GetConfig()
    key := config.GetString("hmac.key")
    mac := hmac.New(sha256.New,[]byte(key))
    mac.Write([]byte(storePhone))
    expMAC := hex.EncodeToString(mac.Sum(nil))
    if expMAC != MAC {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Store id authentication failed."})
        return
    }

    var storeName string
    if store,err := models.GetStoreByPhone(storePhone); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    } else{
        storeName=store.Name
    }
    curTime := time.Now().Local().Unix()
    record := models.Record{
        UserId: userPhone.(string),
        StoreId: storePhone,
        Time: curTime,
    }
    if _,err := record.Record(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    result := forms.Record{
        UserId: record.UserId,
        StoreId: record.StoreId,
        Time: record.Time,
        StoreName: storeName,
    }
    c.JSON(200, result)
}
