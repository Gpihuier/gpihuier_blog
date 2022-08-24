package initialize

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Gpihuier/gpihuier_blog/global"
	"github.com/Gpihuier/gpihuier_blog/route"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunSever() {
	initGlobal()

	routeGroup := route.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.Server.Addr)
	service := initServer(address, routeGroup)

	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`this is for Gpihuier
**        ****        **       ***********      **      **
 **      **  **      **        **               **      **
  **    **    **    **         **  *******      **********
   **  **      **  **          **       **      **      **
    ****        ****           ***********      **      **
访问地址：http://localhost%s
`, address)

	global.LOG.Error(service.ListenAndServe().Error())

}

func initGlobal() {
	global.VP = Viper() // 初始化Viper
	global.DB = GormMysql()
	global.CACHE_DRIVE = RedisDrive()
	global.LOG = Zap()
	zap.ReplaceGlobals(global.LOG) // 替换全局记录器 后续 可以使用 zap.L()
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,          // 监听的 TCP 地址
		Handler:        router,           // Handler：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
		ReadTimeout:    60 * time.Second, // ReadTimeout：允许读取的最大时间
		WriteTimeout:   60 * time.Second, // WriteTimeout：允许写入的最大时间
		MaxHeaderBytes: 1 << 20,          // MaxHeaderBytes：请求头的最大字节数
	}
}
