package models

import "time"

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

type Postsss struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}

type Commentsss struct {
	ID        string    `json:"id"`
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrEmailValidate = "invalid email"
)
