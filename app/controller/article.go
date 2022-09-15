package controller

import (
	"errors"
	"io"
	"strconv"

	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/app/server"
	"github.com/Gpihuier/gpihuier_blog/app/validate"
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/gin-gonic/gin"
)

type Article struct{}

// List 列表
// @method: GET
// @route: /api/article/list
func (a *Article) List(c *gin.Context) {
	var req request.PageInfo
	if err := c.ShouldBindQuery(&req); err != nil {
		if !errors.Is(err, io.EOF) {
			utils.FailWithMessage(err.Error(), c)
			return
		}
	}
	if err := validate.Validate.Common.ValidatePageInfo(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	res, err := server.Server.Article.List(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithData(res, "获取成功", c)
}

// Read 获取单个列表
// @method: GET
// @route: /api/article/list/:id
func (a *Article) Read(c *gin.Context) {
	id := c.Param("id")
	uint64Id, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		utils.FailWithMessage("请输入正整数", c)
		return
	}
	res, err := server.Server.Article.Read(uint64Id)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithData(*res, "获取成功", c)
}

// Create 创建
// @method: POST
// @route: /api/article/save
func (a *Article) Create(c *gin.Context) {
	var req request.ArticleSave
	if err := c.ShouldBindJSON(&req); err != nil {
		if !errors.Is(err, io.EOF) { // 排除io.EOF 错误
			utils.FailWithMessage(err.Error(), c)
			return
		}
	}
	if err := validate.Validate.Article.ArticleSaveValidate(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	// 新增数据
	if err := server.Server.Article.Create(&req, c); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithMessage("新增成功", c)
}

// Update 更新
// @method: PUT
// @route: /api/article/save/:id
func (a *Article) Update(c *gin.Context) {
	var req request.ArticleSave
	if err := c.ShouldBindJSON(&req); err != nil {
		if !errors.Is(err, io.EOF) { // 排除io.EOF 错误
			utils.FailWithMessage(err.Error(), c)
			return
		}
	}
	if err := validate.Validate.Article.ArticleSaveValidate(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	// 获取ID
	id := c.Param("id")
	uint64Id, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		utils.FailWithMessage("请输入正整数", c)
		return
	}
	// 更新数据
	if err = server.Server.Article.Update(uint64Id, &req, c); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithMessage("更新成功", c)
}

// Delete 删除
// @method: DELETE
// @route: /api/article/delete/:id
func (a *Article) Delete(c *gin.Context) {
	id := c.Param("id")
	uint64Id, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		utils.FailWithMessage("请输入正整数", c)
		return
	}
	if err = server.Server.Article.Delete(uint64Id); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithMessage("删除成功", c)
}
