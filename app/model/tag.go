package model

import (
	"errors"
	"fmt"
	"github.com/Gpihuier/gpihuier_blog/app/request"
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

func (t *Tag) Create(req *request.TagSave) error {
	// 判断是否已经有重复的标签名称
	isHas := global.DB.Where("tag_name = ?", req.TagName).First(&t).Error
	if errors.Is(isHas, gorm.ErrRecordNotFound) {
		// 新增数据
		t.TagName = req.TagName
		t.TagColor = req.TagColor
		return global.DB.Create(&t).Error
	}
	return errors.New("标签名称已被使用")
}

func (t *Tag) Update(id uint64, req *request.TagSave) error {
	var res Tag
	isExist := global.DB.First(&res, id).Error
	if errors.Is(isExist, gorm.ErrRecordNotFound) {
		return errors.New("没有找到该标签")
	}
	fmt.Println(res)
	isHas := global.DB.Where("id <> ? AND tag_name = ?", id, req.TagName).First(&res).Error
	fmt.Println(res)
	if errors.Is(isHas, gorm.ErrRecordNotFound) {
		res.TagName = req.TagName
		res.TagColor = req.TagColor
		fmt.Println(res)
		return global.DB.Save(&res).Error
	}
	return errors.New("标签名称已被使用")
}
