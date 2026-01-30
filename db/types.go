package db

import (
	"time"
)

type CategoryDTO struct {
	ID       int32  `json:"id"`
	UserID   int32  `json:"user_id"`
	Name     string `json:"name"`
	Priority int32  `json:"priority"`
}

type ItemDTO struct {
	ID          int32     `json:"id"`
	CategoryID  int32     `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserDTO struct {
	ID           int32  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type CategoryRequest struct {
	Name     string `json:"name"`
	Priority int32  `json:"priority"`
}

type UpdateItemDoneStatusRequest struct {
	Done bool `json:"done"`
}
