package model

import "github.com/Gpihuier/gpihuier_blog/global"

type ArticleContent struct {
	BaseModel
	ArticleId uint64 `gorm:"not null;default:0;comment:文章ID;index:uk_article_id,unique,comment:唯一索引文章ID"`
	Content   string `gorm:"type:text;comment:文章正文"`
}

func (a *ArticleContent) RegisterTable() error {
	if !global.DB.Migrator().HasTable(&a) {
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='文章内容表'").AutoMigrate(&a); err != nil {
			return err
		}
	}
	return nil
}
