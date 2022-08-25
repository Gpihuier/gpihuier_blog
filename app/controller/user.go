package controller

import (
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/app/validate"
	"github.com/Gpihuier/gpihuier_blog/utils"
	"github.com/gin-gonic/gin"
)

type User struct{}

// RegisterUser 注册用户
func (u *User) RegisterUser(c *gin.Context) {
	var req request.RegisterUser
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	// 验证数据
	if err := validate.Validate.User.RegisterUserValidate(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}

}

func (u *User) Login(c *gin.Context) {
	return
}
