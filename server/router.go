package server

import (
    "time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/controllers"
	"github.com/Rhodanthe1116/go-gin-boilerplate/middlewares"
)

func NewRouter(service string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
    router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true //origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

    store := new(controllers.StoreController)
    user := new(controllers.UserController)
    record := new(controllers.RecordController)

    if service=="store" || service=="all" {
        router.POST("/auth/store/signup", store.Signup)
        router.POST("/auth/store/login", store.Login)
    }
    if service == "user" || service=="all" {
        router.POST("/auth/user/signup", user.Signup)
        router.POST("/auth/user/login", user.Login)
    }

    router.Use(middlewares.Authz())
    if service == "store" || service == "all" {
        router.GET("/auth/store/profile", store.Profile)
    }
    if service == "record" || service == "all" {
        router.POST("/records", record.Record)
    }

	return router

}
