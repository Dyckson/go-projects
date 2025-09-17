package domain

import "time"

type User struct {
	UUID      string    `json:"uuid" ksql:"uuid"`
	Name      string    `json:"name" ksql:"name"`
	Email     string    `json:"email" ksql:"email"`
	CreatedAt time.Time `json:"created_at" ksql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" ksql:"updated_at"`
	IsActive  bool      `json:"is_active" ksql:"is_active"`
}

type UserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
