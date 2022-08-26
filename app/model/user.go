package model

import (
	"errors"
	"time"

	"github.com/Gpihuier/gpihuier_blog/global"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Nickname      string    `gorm:"size:20;not null;default:'';comment:用户昵称"`
	Username      string    `gorm:"size:20;uniqueIndex;not null;default:'';comment:用户名/账号"`
	Password      string    `gorm:"size:128;not null;default:'';comment:密码"`
	Avatar        string    `gorm:"size:255;not null;default:'';comment:头像"`
	Email         string    `gorm:"not null;default:'';comment:邮箱"`
	Description   string    `gorm:"size:255;not null;default:'';comment:个人简介"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间"`
	Status        uint8     `gorm:"type:tinyint(1);not null;default:0;comment:登录状态 0 禁止登录 1 正常"`
}

// RegisterTable register user table
func (u *User) RegisterTable() error {
	if !global.DB.Migrator().HasTable(&u) {
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='用户表'").AutoMigrate(&u); err != nil {
			return err
		}
	}
	return nil
}

// JudgeRegisterUserRepeat 检查用户信息是否重复
func (u *User) JudgeRegisterUserRepeat() error {
	var result User
	isHasUsername := global.DB.Where("username = ?", u.Username).First(&result).Error
	if !errors.Is(isHasUsername, gorm.ErrRecordNotFound) {
		return errors.New("用户名已经被注册")
	}
	isHasEmail := global.DB.Where("username = ?", u.Email).First(&result).Error
	if !errors.Is(isHasEmail, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已经被注册")
	}
	return nil
}

func (u *User) RegisterUser() error {
	if err := u.JudgeRegisterUserRepeat(); err != nil {
		return err
	}
	return global.DB.Create(&u).Error
}

// HasUsername 是否存在该用户
func (u *User) HasUsername(username string) (User, error) {
	var result User
	res := global.DB.Where("username = ?", username).First(&result).Error
	if errors.Is(res, gorm.ErrRecordNotFound) {
		return result, errors.New("未找到该用户")
	}
	return result, nil
}

func (u *User) UpdateLastLoginTime(username string) error {
	err := global.DB.Model(&u).Where("username = ?", username).Update("last_login_time", time.Now()).Error
	return err
}
