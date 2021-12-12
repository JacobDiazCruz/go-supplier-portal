package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(ctx *gin.Context) (string, error) {
	cookie, err := ctx.Cookie("token")

	// if token exists
	if err != nil {
		if err == http.ErrNoCookie {
			return "Unauthorized", nil
		}
		return "Bad Request", nil
	}

	tokenStr := cookie
	claims := &Claims{}

	// validate with claims
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "Unauthorized", nil
		}
		return "Bad Request", nil
	}
	if !tkn.Valid {
		return "Unauthorized", nil
	}
	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	fmt.Println("im here token")
	return cookie, nil
	// ctx.JSON(http.StatusOK, gin.H{"msg": "Sucess", "data": claims.Username})
}
