package model

import "github.com/Gpihuier/gpihuier_blog/global"

type Tag struct {
	BaseModel
	TagName string `gorm:"size:30;not null;default:'';comment:标签名称"`
	Color   string `gorm:"size:30;not null;default:'';comment:标签颜色"`
}

func (t *Tag) RegisterTable() error {
	if !global.DB.Migrator().HasTable(&t) {
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='标签表'").AutoMigrate(&t); err != nil {
			return err
		}
	}
	return nil
}
