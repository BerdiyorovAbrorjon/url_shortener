package domain

import (
	"context"
	"time"
)

// db model
type User struct {
	ID             int32     `json:"id"`
	Name           string    `json:"name"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignupResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

type UserResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// API request
type ListUserUrlsRequest struct {
	Limit  int32 `json:"limit" binding:"required,min=1"`
	Offset int32 `json:"offset"`
}

// API response
type ListUserUrlsResponse struct {
	UserID int32 `json:"user_id"`
	Urls   []Url `json:"urls"`
}

type UserRepository interface {
	Create(ctx context.Context, email, hashedPassword, name string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	ListUserUrls(ctx context.Context, userId, limit, offset int32) ([]Url, error)
}

type UserService interface {
	CreateUser(ctx context.Context, email, password, name string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListUserUrls(ctx context.Context, userId, limit, offset int32) ([]Url, error)
}
