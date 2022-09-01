package model

import "github.com/Gpihuier/gpihuier_blog/global"

type ArticleTag struct {
	BaseModel
	ArticleId uint64 `gorm:"not null;default:0;comment:文章ID;index:idx_article_id,unique,comment:普通索引文章id"`
	TagId     uint64 `gorm:"not null;default:0;comment:标签ID;index:idx_tag_id,unique,comment:普通索引标签id"`
}

func (a *ArticleTag) RegisterTable() error {
	if !global.DB.Migrator().HasTable(&a) {
		if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='多对多文章标签中间表'").AutoMigrate(&a); err != nil {
			return err
		}
	}
	return nil
}
