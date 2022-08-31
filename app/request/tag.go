package request

// TagSave 标签新增编辑
type TagSave struct {
	TagName string `json:"tag_name" form:"tag_name" validate:"required,max=30,min=2" alias:"标签名称"`
	Color   string `json:"color" form:"color" validate:"required,rgb" alias:"标签颜色"`
}
