package server

import (
	"errors"
	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/app/response"
	"github.com/Gpihuier/gpihuier_blog/global"
	"github.com/Gpihuier/gpihuier_blog/utils"
	"strconv"
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

func (t *Tag) List(req *request.TagList) (*response.ListData, error) {
	var err error
	var listData response.ListData
	var total int64
	var list []response.TagResponse

	paging, _ := strconv.Atoi(req.Paging)
	db := global.DB.Model(&model.Tag{})

	db.Count(&total)

	if paging > 0 { // 分页返回数据
		page, _ := strconv.Atoi(req.Page)
		if page <= 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(req.PageSize)
		if pageSize <= 0 {
			pageSize = 10
		}
		limit, offset := utils.GetPageData(page, pageSize)
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	if len(list) > 0 {
		for k, v := range list {
			createTime, _ := utils.DateToTime(v.CreateTime, utils.RFC3339Milli)
			list[k].CreateTime = createTime.Format(utils.DEFAULT_YMD)
		}
	}
	listData.Total = total
	listData.List = list
	return &listData, err
}
