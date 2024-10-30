package models

import (
	"time"
)

type User struct {
	ID           int64      `json:"id" db:"id"`
	FirstName    string     `json:"first_name" db:"first_name" validate:"required,min=2,max=100"`
	LastName     string     `json:"last_name" db:"last_name" validate:"required,min=2,max=100"`
	Password     string     `json:"password" db:"password" validate:"required,min=2,max=100"`
	Email        string     `json:"email" db:"email" validate:"email,required"`
	Token        *string    `json:"token" db:"token"`
	RefreshToken *string    `json:"refresh_token" db:"refresh_token"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

// validate eq=ADMIN|eq=USER enum example
