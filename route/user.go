package route

import (
	"github.com/Gpihuier/gpihuier_blog/app/controller"

	"github.com/gin-gonic/gin"
)

type UserRouters struct{}

func (u *UserRouters) Group(r *gin.RouterGroup) {
	apiUser := r.Group("api/user")
	{
		apiUser.POST("register", controller.Controller.User.RegisterUser)
		apiUser.POST("login", controller.Controller.User.Login)
	}
}
