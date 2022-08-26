package server

import (
	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/utils"
	"strings"
)

type User struct{}

func (u *User) RegisterUser(req *request.RegisterUser) error {
	var userModel model.User
	userModel.Nickname = strings.TrimSpace(req.Nickname)
	userModel.Username = strings.TrimSpace(req.Username)
	userModel.Password = utils.BcryptHash(req.Password)
	userModel.Avatar = req.Avatar
	userModel.Email = strings.TrimSpace(req.Email)
	userModel.Description = req.Description
	return userModel.RegisterUser()
}
