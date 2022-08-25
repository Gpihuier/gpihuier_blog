package model

import (
	"github.com/Gpihuier/gpihuier_blog/global"
)

type User struct {
	BaseModel
	Nickname    string `gorm:"size:20;not null;default:'';comment:用户昵称"`
	Username    string `gorm:"size:20;uniqueIndex;not null;default:'';comment:用户名/账号"`
	Password    string `gorm:"size:128;not null;default:'';comment:密码"`
	Avatar      string `gorm:"not null;default:'';comment:头像"`
	Email       string `gorm:"not null;default:'';comment:邮箱"`
	Description string `gorm:"not null;default:'';comment:个人简介"`
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


func (u *User) RegisterUser() error {
	return nil
}
