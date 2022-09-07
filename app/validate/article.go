package validate

import (
	"errors"

	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/global"
	"github.com/Gpihuier/gpihuier_blog/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Article struct{}

func (a *Article) ArticleSaveValidate(req *request.ArticleSave) error {
	validate := validator.New()
	trans := utils.ValidatorTrainInit(validate)
	// 注册自定义验证方法
	if err := validate.RegisterValidation("isCorrectCategoryId", a.isCorrectCategoryId); err != nil {
		return err
	}
	// 注册自定义翻译
	if err := validate.RegisterTranslation("isCorrectCategoryId", trans, func(ut ut.Translator) error {
		return ut.Add("isCorrectCategoryId", "{0}不存在", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isCorrectCategoryId", fe.Field())
		return t
	}); err != nil {
		return err
	}
	if err := validate.Struct(req); err != nil {
		for _, validateErr := range err.(validator.ValidationErrors) {
			return errors.New(validateErr.Translate(trans))
		}
	}
	return nil
}

func (a *Article) isCorrectCategoryId(fl validator.FieldLevel) bool {
	categoryId := fl.Field().String()
	if err := global.DB.Take(&model.Category{}, categoryId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
