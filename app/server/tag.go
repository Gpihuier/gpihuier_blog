package server

import (
	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"strings"
)

type Tag struct{}

func (t *Tag) Create(req *request.TagSave) error {
	req.TagName = strings.TrimSpace(req.TagName)
	req.TagColor = strings.TrimSpace(req.TagColor)
	if err := model.Model.Tag.Create(req); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Update(id uint64, req *request.TagSave) error {
	req.TagName = strings.TrimSpace(req.TagName)
	req.TagColor = strings.TrimSpace(req.TagColor)
	if err := model.Model.Tag.Update(id, req); err != nil {

		return err
	}
	return nil
}
