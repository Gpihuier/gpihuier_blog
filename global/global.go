package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Gpihuier/gpihuier_blog/config"
)

// global 定义全局变量
var (
	VP          *viper.Viper
	DB          *gorm.DB
	CACHE_DRIVE *redis.Client
	LOG         *zap.Logger
	CONFIG      config.Enter
)
