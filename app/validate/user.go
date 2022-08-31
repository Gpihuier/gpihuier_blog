package validate

import (
	"errors"

	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/go-playground/validator/v10"
)

type User struct{}

// RegisterUserValidate 注册用户验证
func (u *User) RegisterUserValidate(req *request.RegisterUser) error {
	validate := validator.New()
	validate.RegisterStructValidation(u.ConfirmPasswordValidate, req)
	trans := utils.ValidatorTrainInit(validate)
	if err := validate.Struct(req); err != nil {
		for _, validateErr := range err.(validator.ValidationErrors) {
			return errors.New(validateErr.Translate(trans))
		}
	}
	return nil
}

func (u *User) ConfirmPasswordValidate(sl validator.StructLevel) {
	req := sl.Current().Interface().(request.RegisterUser)
	if req.Password != req.ConfirmPassword {
		sl.ReportError(req.ConfirmPassword, "confirm_password", "ConfirmPassword", "eqfield", "password")
	}
}

// LoginValidate 登录用户验证
func (u *User) LoginValidate(req *request.Login) error {
	validate := validator.New()
	trans := utils.ValidatorTrainInit(validate)
	if err := validate.Struct(req); err != nil {
		for _, validateErr := range err.(validator.ValidationErrors) {
			return errors.New(validateErr.Translate(trans))
		}
	}
	return nil
}
