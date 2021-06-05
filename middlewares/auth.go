package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/Rhodanthe1116/go-gin-boilerplate/auth"
    "github.com/Rhodanthe1116/go-gin-boilerplate/config"
)

func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}
		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}
        config := config.GetConfig()
		jwtWrapper := auth.JwtWrapper{
            SecretKey:       config.GetString("jwt.secret"),
			Issuer:    "AuthService",
		}
		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		c.Set("phone", claims.Phone)
		c.Next()
	}
}

