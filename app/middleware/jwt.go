package middleware

import (
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			utils.FailWithMessage("请先登录", c)
			c.Abort()
			return
		}
		jwt := utils.NewJwtSecret()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			utils.FailWithMessage("登录失效", c)
			c.Abort()
			return
		}
		c.Set("userinfo", claims)
		c.Next()
	}
}
