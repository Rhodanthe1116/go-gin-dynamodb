package server

import (
	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	hello := new(controllers.HelloController)
	health := new(controllers.HealthController)

	router.GET("/", hello.Home)

	router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.POST("/signup", user.Signup)
			userGroup.GET("/:id", user.Retrieve)
		}
	}
	return router

}
