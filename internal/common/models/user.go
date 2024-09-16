// internal/common/models/user.go
package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Email       string    `db:"email" json:"email"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	Avatar      *string   `db:"avatar" json:"avatar"`
	Password    string    `db:"password" json:"password"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	Role        string    `db:"role" json:"role"`
	RegisterBy  string    `db:"register_by" json:"register_by"`
	CreatedAt   string    `db:"created_at" json:"created_at"`
	UpdatedAt   string    `db:"updated_at" json:"updated_at"`
}
