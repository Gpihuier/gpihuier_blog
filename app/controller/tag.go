// Package controller 标签的增删改查
package controller

import (
	"errors"
	"fmt"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/app/server"
	"github.com/Gpihuier/gpihuier_blog/app/validate"
	"github.com/Gpihuier/gpihuier_blog/utils"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

// List 列表
// @method: GET
// @route: /api/tag/list
func (t *Tag) List(c *gin.Context) {
	var req request.TagSave
	if err := c.ShouldBindJSON(&req); err != nil {
		if !errors.Is(err, io.EOF) { // 排除io.EOF 错误
			utils.FailWithMessage(err.Error(), c)
			return
		}
	}
	if err := validate.Validate.Tag.TagSaveValidate(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(123)
}

// Read 编辑查看详细
// @method: GET
// @route: /api/tag/list/:id
func (t *Tag) Read(c *gin.Context) {
	fmt.Println(c.Param("id"))
}

// Create 新增标签
// @method: POST
// @route: /api/tag/save
func (t *Tag) Create(c *gin.Context) {
	var req request.TagSave
	if err := c.ShouldBindJSON(&req); err != nil {
		if !errors.Is(err, io.EOF) { // 排除io.EOF 错误
			utils.FailWithMessage(err.Error(), c)
			return
		}
	}
	if err := validate.Validate.Tag.TagSaveValidate(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	if err := server.Server.Tag.Create(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithMessage("新增成功", c)
}

// Update 编辑标签
// @method: PUT
// @route: /api/tag/save/:id
func (t *Tag) Update(c *gin.Context) {
	var req request.TagSave
	if err := c.ShouldBindJSON(&req); err != nil {
		if !errors.Is(err, io.EOF) { // 排除io.EOF 错误
			utils.FailWithMessage(err.Error(), c)
			return
		}
	}
	if err := validate.Validate.Tag.TagSaveValidate(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	// 获取uri参数
	id := c.Param("id")
	uint64Id, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		utils.FailWithMessage("请输入整数", c)
		return
	}
	if err = server.Server.Tag.Update(uint64Id, &req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	utils.SuccessWithMessage("更新成功", c)
}

// Delete 删除标签
// @method: DELETE
// @route: /api/tag/delete/:id
func (t *Tag) Delete(c *gin.Context) {

}
