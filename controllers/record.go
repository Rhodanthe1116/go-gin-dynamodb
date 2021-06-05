package controllers

import (
    "time"

	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/forms"
	"net/http"
)

type RecordController struct{}

func (h RecordController) Record(c *gin.Context) {
    var Payload forms.RecordStoreId
    if err := c.ShouldBindJSON(&Payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    phone,_ := c.Get("phone")
    curTime := time.Now().Local().Unix()
    // TODO: user=db.get(phone);
    // TODO: db.insert(Payload.StoreId, user.id, curTime)
    // c.String(200, "Success")
    var Response forms.RecordRecord
    Response.StoreId=Payload.StoreId
    Response.UserId="998244353"
    Response.Time=curTime
    Response.Phone=phone.(string)
    c.JSON(200, Response)
}
