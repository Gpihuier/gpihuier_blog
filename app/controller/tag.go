// Package controller 标签的增删改查
package controller

import "github.com/gin-gonic/gin"

type Tag struct{}

// List 列表
// @method: GET
// @route: /api/tag/list
func (t *Tag) List(c *gin.Context) {

}

// Read 编辑查看详细
// @method: GET
// @route: /api/tag/list/:id
func (t *Tag) Read(c *gin.Context) {

}

// Create 新增标签
// @method: POST
// @route: /api/tag/save
func (t *Tag) Create(c *gin.Context) {

}

// Update 编辑标签
// @method: PUT
// @route: /api/tag/save/:id
func (t *Tag) Update(c *gin.Context) {

}

// Delete 删除标签
// @method: DELETE
// @route: /api/tag/delete/:id
func (t *Tag) Delete(c *gin.Context) {

}
