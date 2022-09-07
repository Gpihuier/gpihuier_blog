package model

import "github.com/Gpihuier/gpihuier_blog/global"

type Article struct {
	BaseModel
	Title          string         `gorm:"size:50;not null;default:'';comment:文章标题"`
	AuthorId       uint64         `gorm:"not null;default:0;comment:作者ID"`
	CategoryId     uint64         `gorm:"not null;default:0;comment:类别ID"`
	Status         uint8          `gorm:"type:tinyint(1);not null;default:0;comment:发布状态 0 草稿/未发布 1 已发布"`
	IsTop          uint8          `gorm:"type:tinyint(1);not null;default:0;comment:是否置顶 0 否 1是"`
	ArticleContent ArticleContent // has one
}

type ArticleContent struct {
	BaseModel
	ArticleId uint64 `gorm:"not null;default:0;comment:文章ID;index:uk_article_id,unique,comment:唯一索引文章ID"`
	Content   string `gorm:"type:text;comment:文章正文"`
}

func (a *Article) RegisterTable() error {
	if !global.DB.Migrator().HasTable(&a) {
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='文章表'").AutoMigrate(&a); err != nil {
			return err
		}
		// 创建文章正文表
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='文章正文表'").AutoMigrate(&ArticleContent{}); err != nil {
			return err
		}
	}
	return nil
}
