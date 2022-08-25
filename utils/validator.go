package utils

import (
	"reflect"
	"strings"

	"github.com/Gpihuier/gpihuier_blog/global"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
)

// ValidatorTrainInit 初始化验证器 注册中文翻译器
func ValidatorTrainInit(validate *validator.Validate) ut.Translator {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	if err := zhTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		global.LOG.Error("validate init error", zap.Error(err))
	}
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string { // 增加字段别名
		name := strings.SplitN(fld.Tag.Get("alias"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return trans
}
