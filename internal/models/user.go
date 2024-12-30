package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"Email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUserValidate struct {
	Email string `validate:"required,email"`
}

var (
	ErrEmailValidate = "invalid email"
)
