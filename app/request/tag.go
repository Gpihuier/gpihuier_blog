package request

// TagSave 标签新增编辑
type TagSave struct {
	TagName  string `json:"tag_name" validate:"required,max=30,min=2" alias:"标签名称"`
	TagColor string `json:"tag_color" validate:"required,rgb" alias:"标签颜色"`
}

// TagList 列表
type TagList struct {
	PageInfo
	TagName string `json:"tag_name" validate:"omitempty,max=30,min=2" alias:"标签名称"`
}
