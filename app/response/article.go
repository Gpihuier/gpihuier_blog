package response

type Article struct {
	Id             uint64         `json:"id"`
	Title          string         `json:"title"`
	CategoryId     uint64         `json:"category_id"`
	Status         uint8          `json:"status"`
	ArticleContent ArticleContent `json:"article_content"`
}

type ArticleContent struct {
	ArticleId uint64 `json:"article_id"`
	Content   string `json:"content"`
}

type ArticleList struct {
	Id            uint64 `json:"id"`
	CreateTime    string `json:"create_time"`
	Title         string `json:"title"`
	AuthorId      uint64 `json:"author_id"`
	CategoryId    uint64 `json:"category_id"`
	IsTop         uint8  `json:"is_top"`
	CategoryTitle string `json:"category_title"`
}
