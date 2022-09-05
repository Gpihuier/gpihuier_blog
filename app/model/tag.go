package model

import (
	"errors"

	"github.com/Gpihuier/gpihuier_blog/global"

	"gorm.io/gorm"
)

type Tag struct {
	BaseModel
	TagName  string `gorm:"size:30;not null;default:'';comment:标签名称;index:uk_tag_name,unique,comment:唯一索引标签名称"`
	TagColor string `gorm:"size:30;not null;default:'';comment:标签颜色"`
}

func (t *Tag) RegisterTable() error {
	if !global.DB.Migrator().HasTable(&t) {
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='标签表'").AutoMigrate(&t); err != nil {
			return err
		}
	}
	return nil
}

// IsHasTagName 判断是否存在标签名称
func (t *Tag) IsHasTagName() bool {
	var res Tag
	var isHas error
	if t.ID > 0 {
		isHas = global.DB.Where("id <> ? AND tag_name = ?", t.ID, t.TagName).First(&res).Error
	} else {
		isHas = global.DB.Where("tag_name = ?", t.TagName).First(&res).Error
	}
	if isHas == nil {
		return true
	} else {
		return false
	}
}

// IsExist 判断更新是否存在
func (t *Tag) IsExist() (*Tag, error) {
	var res Tag
	if errors.Is(global.DB.First(&res, t.ID).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("没有找到该标签")
	}
	return &res, nil
}

func (t *Tag) Create() error {
	return global.DB.Create(&t).Error
}

func (t *Tag) Update() error {
	return global.DB.Where("id = ?", t.ID).First(&Tag{}).Updates(&t).Error
}
