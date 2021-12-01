package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

type TokenIdentity struct {
	Username string
}

func SignToken(tk TokenIdentity) string {
	fmt.Println(tk)
	// err := tk.Username

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
