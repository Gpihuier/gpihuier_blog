package response

// ListData 分页返回数据
type ListData struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}
