package server

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Gpihuier/gpihuier_blog/app/model"
	"github.com/Gpihuier/gpihuier_blog/app/request"
	"github.com/Gpihuier/gpihuier_blog/global"
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (a *Article) Update(id uint64, req *request.ArticleSave, c *gin.Context) error {
	var res model.Article
	isHas := global.DB.Preload("ArticleContent").First(&res, id)
	if errors.Is(isHas.Error, gorm.ErrRecordNotFound) {
		return ErrNotExist
	}
	// 获取参数
	categoryId, _ := strconv.ParseUint(req.CategoryId, 0, 64)
	status, _ := strconv.ParseUint(req.Status, 0, 64)
	token := c.GetHeader("Authorization")
	claims, err := utils.NewJwtSecret().ParseToken(token)
	if err != nil {
		return err
	}
	// 开启事务
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 先更新主表
		res.Title = strings.TrimSpace(req.Title)
		res.AuthorId = claims.BaseClaims.ID
		res.CategoryId = categoryId
		res.Status = uint8(status)
		if err := tx.Where("id = ?", id).First(&model.Article{}).Updates(&res).Error; err != nil {
			return err
		}
		// 在更新附表
		content := model.ArticleContent{
			Content: req.Content,
		}
		oldContent := tx.Where("article_id = ?", id).First(&model.ArticleContent{})
		if errors.Is(oldContent.Error, gorm.ErrRecordNotFound) {
			return errors.New("数据不完整, 请删除后重试")
		}
		if err := oldContent.Updates(&content).Error; err != nil {
			return err
		}
		return nil // 返回 nil 提交事务
	})
}

func (a *Article) Read(id uint64) (*map[string]any, error) {
	data := make(map[string]any)
	err := global.DB.Model(&model.Article{}).
		Where("id = ?", id).
		Preload("ArticleContent").
		Omit("UpdateTime").
		First(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("没有找到该文章")
	}

	return &data, nil
}
