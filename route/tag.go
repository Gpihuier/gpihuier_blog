package route

import (
	"github.com/Gpihuier/gpihuier_blog/app/controller"

	"github.com/gin-gonic/gin"
)

type TagRouters struct{}

func (t *TagRouters) Group(r *gin.RouterGroup) {
	tag := r.Group("api/tag")
	{
		tag.POST("save", controller.Controller.Tag.Create)
		tag.PUT("save/:id", controller.Controller.Tag.Update)
	}
	{
		tag.GET("list", controller.Controller.Tag.List)
		tag.GET("list/:id", controller.Controller.Tag.Read)
	}
	{
		tag.DELETE("delete/:id", controller.Controller.Tag.Delete)
	}
}
