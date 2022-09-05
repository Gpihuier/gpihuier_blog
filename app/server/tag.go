package server

import (
	"errors"
	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"strings"
)

var ErrHasTagName = errors.New("标签名称已经存在")

type Tag struct{}

func (t *Tag) Create(req *request.TagSave) error {
	var res model.Tag
	res.TagName = strings.TrimSpace(req.TagName)
	res.TagColor = req.TagColor
	if res.IsHasTagName() == true {
		return ErrHasTagName
	}
	return res.Create()
}

func (t *Tag) Update(id uint64, req *request.TagSave) error {
	var res model.Tag
	res.ID = id
	exist, err := res.IsExist()
	if err != nil {
		return err
	}
	exist.TagName = strings.TrimSpace(req.TagName)
	exist.TagColor = req.TagColor
	if exist.IsHasTagName() == true {
		return ErrHasTagName
	}
	return exist.Update()
}
