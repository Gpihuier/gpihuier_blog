package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"time"

	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GormMysql() *gorm.DB {
	config := global.CONFIG.Mysql

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DSN(),
		DefaultStringSize:         191,  // 字符串的默认长度
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		SkipInitializeWithVersion: true, // // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Prefix, // 表名前缀，`User` 的表名应该是 `prefix_users`
			SingularTable: true,          // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `prefix_user`
		},
		SkipDefaultTransaction:                   true, // 跳过默认事务
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建数据库外键约束
		Logger: logger.New( // 日志
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // 禁用彩色打印
			},
		),
	})

	if err != nil {
		panic(fmt.Errorf("链接数据库出现错误:%s", err))
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db

}

// RegisterTable 自动迁移
func RegisterTable() {
	// TODO 函数式错误编程
	var err error
	enter := model.Model
	if err = enter.User.RegisterTable(); err != nil {
		global.LOG.Error("register User Table failed", zap.Error(err))
		panic(fmt.Errorf("register user table failed: %v", err))
	}
	if err = enter.Article.RegisterTable(); err != nil {
		global.LOG.Error("register Article Table failed", zap.Error(err))
		panic(fmt.Errorf("register Article table failed: %v", err))
	}
	if err = enter.Category.RegisterTable(); err != nil {
		global.LOG.Error("register Category Table failed", zap.Error(err))
		panic(fmt.Errorf("register Category table failed: %v", err))
	}
	if err = enter.Tag.RegisterTable(); err != nil {
		global.LOG.Error("register tagTable failed", zap.Error(err))
		panic(fmt.Errorf("register tag table failed: %v", err))
	}
	if err = enter.ArticleTag.RegisterTable(); err != nil {
		global.LOG.Error("register ArticleTag Table failed", zap.Error(err))
		panic(fmt.Errorf("register ArticleTag table failed: %v", err))
	}

	global.LOG.Info("register tables success!")
}
