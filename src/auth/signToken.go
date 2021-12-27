package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key")

/**
 * @author Jacob
 * @description verify token via google oauth
 * @description register user and sign jwt token
 * @param - token, login type
 * @returns - access token
 */
func SignToken(email string) string {
	// sign jwt
	claims := &Claims{
		Username: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Minute * 5)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	// if jwt signature is invalid
	if err != nil {
		return "Error"
	}

	// if success
	return tokenString
}
