package response

type TagResponse struct {
	Id         uint64 `json:"id"`
	CreateTime string `json:"create_time"`
	TagName    string `json:"tag_name"`
	TagColor   string `json:"tag_color"`
}
