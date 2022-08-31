package request

// PageInfo 分页参数
type PageInfo struct {
	Paging   int `json:"paging" form:"paging" comment:"是否分页 0分页 1不分页"`
	Page     int `json:"page" form:"page" comment:"是否分页 0分页 1不分页"`
	PageSize int `json:"page_size" form:"page_size" comment:"是否分页 0分页 1不分页"`
}
