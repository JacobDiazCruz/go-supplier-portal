package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

/**
 * @author Gio
 * @description verify token via google oauth
 * @description register user and sign jwt token
 * @param - token, login type
 * @returns - access token
 */
func SignToken(tk TokenIdentity) string {
	// add expiration time
	expirationTime := time.Now().Add(time.Minute * 5)

	// sign jwt
	claims := &Claims{
		Username: tk.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
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
