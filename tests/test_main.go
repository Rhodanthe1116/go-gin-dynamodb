package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/Rhodanthe1116/go-gin-boilerplate/controllers"
	"github.com/Rhodanthe1116/go-gin-boilerplate/db"
)

// func Test(t *testing.T) { Testing(t) }

// var _ = Suite(&UserSuite{})

type UserSuite struct {
	config *viper.Viper
	router *gin.Engine
}

// func (s *UserSuite) SetUpTest(c *C) {
// 	config.Init("test")
// 	s.config = config.GetConfig()
// 	s.router = SetupRouter()
// }

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	hello := new(controllers.HelloController)
	health := new(controllers.HealthController)
	router.GET("/", hello.Home)
	router.GET("/health", health.Status)

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

func TestMain(m *testing.M) {
	db.Init()
	SetupRouter()
}
