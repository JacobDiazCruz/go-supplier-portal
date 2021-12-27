package auth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(ctx *gin.Context) *TokenIdentity {
	if len(ctx.Request.Header["Authorization"]) <= 0 {
		fmt.Println("here error token")
		return nil
	}
	token := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]

	fmt.Println(token)
	fmt.Println("here token")

	tokenStr := token
	claims := &Claims{}

	// validate with claims
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil
		}
		return nil
	}
	if !tkn.Valid {
		return nil
	}
	// do something with decoded claims
	// for key, val := range tkn {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	// fmt.Printf("%v\n", claims)
	tk := &TokenIdentity{
		Username: claims.Username,
		Token:    token,
	}
	return tk
}
