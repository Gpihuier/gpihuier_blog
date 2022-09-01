package validate

import (
	"errors"

	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/go-playground/validator/v10"
)

type Tag struct{}

func (t *Tag) TagSaveValidate(req *request.TagSave) error {
	validate := validator.New()
	trans := utils.ValidatorTrainInit(validate)
	if err := validate.Struct(req); err != nil {
		for _, validateErr := range err.(validator.ValidationErrors) {
			return errors.New(validateErr.Translate(trans))
		}
	}
	return nil
}
