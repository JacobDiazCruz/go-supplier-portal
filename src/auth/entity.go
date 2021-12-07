package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenIdentity struct {
	Username string
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Credentials struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Type  string `json:"type"`
	Scope string `json:"scope"`
}
