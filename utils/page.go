package utils

// GetPageData 获取分页数据
func GetPageData(page, pageSize int) (limit, offset int) {
	limit = pageSize
	offset = pageSize * (page - 1)
	return
}
