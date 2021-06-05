package server

import (
	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/controllers"
	"github.com/Rhodanthe1116/go-gin-boilerplate/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

    store := new(controllers.StoreController)
    user := new(controllers.UserController)
    record := new(controllers.RecordController)

    router.POST("/auth/store/signup", store.Signup)
    router.POST("/auth/store/login", store.Login)
    router.POST("/auth/user/signup", user.Signup)
    router.POST("/auth/user/login", user.Login)

    router.Use(middlewares.Authz())
    router.GET("/auth/store/profile", store.Profile)
    router.POST("/records", record.Record)

	return router

}
