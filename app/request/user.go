package request

type RegisterUser struct {
	Nickname        string `json:"nickname" validate:"required,max=10,min=2" alias:"昵称"`
	Username        string `json:"username" validate:"required,alphanum,max=10,min=2" alias:"用户名"`
	Password        string `json:"password" validate:"required,alphanum,max=20,min=6" alias:"密码"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password" alias:"确认密码"`
	Avatar          string `json:"avatar" validate:"omitempty,max=255,min=1" alias:"头像"`
	Email           string `json:"email" validate:"omitempty,email" alias:"邮箱"`
	Description     string `json:"description" validate:"omitempty,max=255" alias:"个人简介"`
}
