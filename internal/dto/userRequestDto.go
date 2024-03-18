package dto

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterDto struct {
	UserLoginDto
	Phone string `json:"phone"`
}
