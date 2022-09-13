package request

// ArticleSave 文章新增修改的保存
type ArticleSave struct {
	Title      string `json:"title" validate:"required,max=50,min=3" alias:"文章标题"`
	CategoryId string `json:"category_id" validate:"required,number,isCorrectCategoryId" alias:"文章分类"`
	Status     string `json:"status" validate:"required,number,oneof=0 1" alias:"文章状态"`
	Content    string `json:"content" validate:"required,max=60000,min=1" alias:"文章正文"`
}

type ArticleList struct {
	PageInfo PageInfo
	Keyword  string `json:"keyword" validate:"omitempty,max=30,min=2" alias:"关键词"`
}
