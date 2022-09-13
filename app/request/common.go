package request

// PageInfo 分页参数
type PageInfo struct {
	Paging   string `form:"paging" validate:"omitempty,number" alias:"是否分页"`
	Page     string `form:"page" validate:"omitempty,number" alias:"当前页码"`
	PageSize string `form:"page_size" validate:"omitempty,number" alias:"分页数量"`
	Keyword  string `form:"keyword" validate:"omitempty" alias:"关键词"`
}
