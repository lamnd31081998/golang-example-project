package customstruct

type RegisterDto struct {
	Username         string `json:"username" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Password         string `json:"password" validate:"required,min=6"`
	Confirm_Password string `json:"confirm_password" validate:"required,min=6,eqfield=Password"`
}

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type TokenInfo struct {
	Token  string
	UserId uint
	User   User
}
