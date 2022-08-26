package route

import (
	"net/http"

	"github.com/Gpihuier/gpihuier_blog/app/middleware"
	"github.com/Gpihuier/gpihuier_blog/global"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	//Router := gin.New()
	//Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	//gin.SetMode("debug")
	Router := gin.Default()

	var HelloRouter = Router.Group("") // 测试路由
	{
		HelloRouter.GET("/hello/:name", func(c *gin.Context) {
			var msg = c.Param("name")
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello " + msg,
			})
		})
	}

	// 基础路由不做鉴权
	PublicRouters := Router.Group("")
	{
		RouterEnter.UserRouters.Group(PublicRouters)
	}

	// 登录鉴权
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JwtMiddleware())
	{

	}

	global.LOG.Info("router register success")
	return Router
}
