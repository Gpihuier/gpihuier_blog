package server

import (
	"github.com/Gpihuier/gpihuier_blog/utils"
	"strconv"
	"strings"

	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/global"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func (a *Article) Create(req *request.ArticleSave, c *gin.Context) error {
	categoryId, _ := strconv.ParseUint(req.CategoryId, 0, 64)
	status, _ := strconv.ParseUint(req.Status, 0, 64)
	token := c.GetHeader("Authorization")
	claims, err := utils.NewJwtSecret().ParseToken(token)
	if err != nil {
		return err
	}
	res := model.Article{
		Title:          strings.TrimSpace(req.Title),
		AuthorId:       claims.BaseClaims.ID,
		CategoryId:     categoryId,
		Status:         uint8(status),
		IsTop:          uint8(0),
		ArticleContent: model.ArticleContent{Content: req.Content},
	}
	// 预加载
	return global.DB.Preload("ArticleContent").Create(&res).Error
}
