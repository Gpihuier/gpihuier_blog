package route

import (
	"github.com/Gpihuier/gpihuier_blog/app/controller"
	"github.com/gin-gonic/gin"
)

type ArticleRouters struct{}

func (a *ArticleRouters) Group(r *gin.RouterGroup) {
	routes := r.Group("api/article")
	{
		routes.POST("save", controller.Controller.Article.Create)
		routes.PUT("save/:id", controller.Controller.Article.Update)
	}
	{
		routes.GET("list", controller.Controller.Article.List)
		routes.GET("list/:id", controller.Controller.Article.Read)
	}
	{
		routes.DELETE("delete/:id", controller.Controller.Article.Delete)
	}
}
