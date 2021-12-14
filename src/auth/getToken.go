package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(ctx *gin.Context) (*TokenIdentity, error) {
	cookie, err := ctx.Cookie("token")

	// if token exists
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, nil
		}
		return nil, nil
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
			return nil, nil
		}
		return nil, nil
	}
	if !tkn.Valid {
		return nil, nil
	}
	// do something with decoded claims
	// for key, val := range tkn {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	// fmt.Printf("%v\n", claims)
	tk := &TokenIdentity{
		Username: claims.Username,
		Token:    cookie,
	}
	return tk, nil
	// ctx.JSON(http.StatusOK, gin.H{"msg": "Sucess", "data": claims.Username})
}
