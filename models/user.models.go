package models

import "time"

type User struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Disable   bool      `json:"disable"`
	CreatedAt time.Time `json:"created_at"`
}
