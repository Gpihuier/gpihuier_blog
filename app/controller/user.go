package controller

import (
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/app/server"
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
	if err := server.Server.User.RegisterUser(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithMessage("注册成功", c)
}

// Login 用户登录
func (u *User) Login(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	if err := validate.Validate.User.LoginValidate(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	res, err := server.Server.User.Login(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithData(res, "登录成功", c)
}
