package response

type Login struct {
	Id            uint64 `json:"id"`
	Nickname      string `json:"nickname"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Email         string `json:"email"`
	Description   string `json:"description"`
	LastLoginTime string `json:"last_login_time"`
	Token         string `json:"token"`
	ExpiresAt     string `json:"expires_at"`
}
