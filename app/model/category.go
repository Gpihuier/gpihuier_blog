package model

import "github.com/Gpihuier/gpihuier_blog/global"

type Category struct {
	BaseModel
	Title string `gorm:"size:20;not null;default:'';comment:类别名称;uniqueIndex:uk_title,comment:唯一索引类别名称"`
}

func (c *Category) RegisterTable() error {
	if !global.DB.Migrator().HasTable(&c) {
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='类别表'").AutoMigrate(&c); err != nil {
			return err
		}
	}
	return nil
}
