package server

import (
	"context"
	"errors"
	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/app/response"
	"github.com/Gpihuier/gpihuier_blog/global"
	"github.com/Gpihuier/gpihuier_blog/utils"
	"go.uber.org/zap"
	"strings"
	"time"
)

type User struct{}

func (u *User) RegisterUser(req *request.RegisterUser) error {
	userModel := model.Model.User
	userModel.Nickname = strings.TrimSpace(req.Nickname)
	userModel.Username = strings.TrimSpace(req.Username)
	userModel.Password = utils.BcryptHash(req.Password)
	userModel.Avatar = req.Avatar
	userModel.Email = strings.TrimSpace(req.Email)
	userModel.Description = req.Description
	return userModel.RegisterUser()
}

func (u *User) Login(req *request.Login) (*response.Login, error) {
	userModel := model.Model.User
	username := strings.TrimSpace(req.Username)
	// 判断是否存在用户
	result, err := userModel.HasUsername(username)
	if err != nil {
		return nil, err
	}
	if result.Status != 1 {
		return nil, errors.New("用户禁止登录")
	}
	// 判断密码
	if !utils.BcryptCheck(req.Password, result.Password) {
		return nil, errors.New("密码错误")
	}
	// 修改上次登录时间
	if err = userModel.UpdateLastLoginTime(username); err != nil {
		return nil, err
	}
	// 签发jwt
	jwt := utils.NewJwtSecret()
	claims := jwt.CreateClaims(request.BaseClaims{
		ID:          result.ID,
		Nickname:    result.Nickname,
		Username:    result.Username,
		Avatar:      result.Avatar,
		Email:       result.Email,
		Description: result.Description,
	})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		global.LOG.Error("用户登录时生成token失败", zap.Error(err))
		return nil, err
	}
	// 保存到redis
	redisJwtExp := time.Second * time.Duration(global.CONFIG.Jwt.ExpiresTime)
	err = global.CACHE_DRIVE.Set(context.Background(), result.Username, token, redisJwtExp).Err()
	if err != nil {
		global.LOG.Error("用户登录时存入token到redis失败", zap.Error(err))
		return nil, err
	}
	res := response.Login{
		Id:            result.ID,
		Nickname:      result.Nickname,
		Username:      result.Username,
		Avatar:        result.Avatar,
		Email:         result.Email,
		Description:   result.Description,
		LastLoginTime: result.LastLoginTime.Format("2006-01-02 15:04:05"),
		Token:         token,
		ExpiresAt:     utils.TimestampToDate(time.Now().Unix()+global.CONFIG.Jwt.ExpiresTime+60, ""),
	}
	return &res, nil
}
