package controller

import (
	"errors"
	"github.com/Gpihuier/gpihuier_blog/app/server"
	"io"

	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/app/validate"
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/gin-gonic/gin"
)

type Article struct{}

// List 列表
// @method: GET
// @route: /api/article/list
func (a *Article) List(c *gin.Context) {

}

// Read 获取单个列表
// @method: GET
// @route: /api/article/list/:id
func (a *Article) Read(c *gin.Context) {

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

}

// Delete 删除
// @method: DELETE
// @route: /api/article/delete/:id
func (a *Article) Delete(c *gin.Context) {

}
