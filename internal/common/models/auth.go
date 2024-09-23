package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	FirstName string    `json:"first_name"`
	LastNamme string    `json:"last_name"`
	ID        uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

type UserClaims struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastNamme string    `json:"last_name"`
}
