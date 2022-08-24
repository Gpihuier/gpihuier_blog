package route

import (
	"net/http"

	"github.com/Gpihuier/gpihuier_blog/global"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	//Router := gin.New()
	//Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	//gin.SetMode("debug")
	Router := gin.Default()

	global.LOG.Info("register swagger handler")

	var HelloRouter = Router.Group("") // 测试路由
	{
		HelloRouter.GET("/hello/:name", func(c *gin.Context) {
			var msg = c.Param("name")
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello " + msg,
			})
		})
	}
	global.LOG.Info("router register success")
	return Router
}
