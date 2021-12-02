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
