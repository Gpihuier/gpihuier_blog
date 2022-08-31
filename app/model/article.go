package model

import "github.com/Gpihuier/gpihuier_blog/global"

type Article struct {
	BaseModel
	Title      string `gorm:"size:30;not null;default:'';comment:文章标题"`
	AuthorId   uint   `gorm:"not null;default:'';comment:作者ID"`
	CategoryId uint   `gorm:"not null;default:'';comment:类别ID"`
	Status     uint8  `gorm:"type:tinyint(1);not null;default:0;comment:发布状态 0 草稿/未发布 1 已发布"`
}

func (a *Article) RegisterTable() error {
	if !global.DB.Migrator().HasTable(&a) {
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='文章表'").AutoMigrate(&a); err != nil {
			return err
		}
	}
	return nil
}
