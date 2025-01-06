package models

import "database/sql"

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Bio        string `json:"bio"`
	ProfilePic string `json:"profile_pic"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

type GetUserResponse struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	ProfilePic string `json:"profile_pic"`
	Bio        string `json:"bio"`
}

type GetUserValidate struct {
	Email string `validate:"required,email"`
}

type GetAllFollowersResponseDB struct {
	Username   string         `json:"username"`
	ProfilePic sql.NullString `json:"profile_pic"`
}

type GetAllFollowersResponse struct {
	Username   string `json:"username"`
	ProfilePic string `json:"profile_pic"`
}

type UserID struct {
	ID int `json:"id"`
}

var (
	ErrEmailValidate = "invalid email"
)
