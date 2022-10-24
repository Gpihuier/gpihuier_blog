package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPageData 获取分页数据
func GetPageData(page, pageSize int) (limit, offset int) {
	limit = pageSize
	offset = pageSize * (page - 1)
	return
}

// Paginate Gorm 分页 https://gorm.cn/zh_CN/docs/scopes.html#pagination
// example: db.Scopes(Paginate(c)).Find(&users)
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
