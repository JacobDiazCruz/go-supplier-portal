package authentication

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

	return cookie, nil
	// ctx.JSON(http.StatusOK, gin.H{"msg": "Sucess", "data": claims.Username})
}
