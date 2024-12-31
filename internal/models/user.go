package models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Bio        string `json:"bio"`
	ProfilePic string `json:"profile_pic"`
	Followers  []int  `json:"followers"`
	Following  []int  `json:"following"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

type GetUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUserValidate struct {
	Email string `validate:"required,email"`
}

type GetAllFollowersRequest struct {
	Username   string `json:"username"`
	ProfilePic string `json:"profile_pic"`
}

var (
	ErrEmailValidate = "invalid email"
)
