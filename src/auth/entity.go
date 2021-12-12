package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type TokenIdentity struct {
	Username string
}
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Type  string `json:"type"`
	Scope string `json:"scope"`
}
