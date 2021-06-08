package controllers

import (
    "time"

	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/forms"
	"github.com/Rhodanthe1116/go-gin-boilerplate/models"
	"net/http"
)

type RecordController struct{}

func (h RecordController) Record(c *gin.Context) {
    var Payload forms.RecordStoreId
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    userPhone,_ := c.Get("phone")

    storePhone := Payload.StoreId
    if _,err := models.GetStoreByPhone(storePhone); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
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
    c.JSON(200, record)
}
