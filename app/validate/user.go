package validate

import (
	"errors"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/go-playground/validator/v10"
)

type User struct{}

func (u *User) RegisterUserValidate(r *request.RegisterUser) error {
	validate := validator.New()
	trans := utils.ValidatorTrainInit(validate)
	if err := validate.Struct(r); err != nil {
		for _, validateErr := range err.(validator.ValidationErrors) {
			return errors.New(validateErr.Translate(trans))
		}
	}
	return nil
}
