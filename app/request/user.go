package request

type RegisterUser struct {
	Nickname string `json:"nickname" validate:"required"`
}
