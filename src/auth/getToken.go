package auth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	getUser "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func GetToken(ctx *gin.Context) *TokenIdentity {
	if len(ctx.Request.Header["Authorization"]) <= 0 {
		fmt.Println("here error token")
		return nil
	}
	token := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]

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

	// get user by email
	fmt.Println(claims.Email)
	fmt.Println("email here123")
	userData := getUser.GetEmail(claims.Email)

	// return response
	tk := &TokenIdentity{
		Username:       userData.Username,
		Email:          userData.Email,
		ThumbnailImage: userData.ThumbnailImage,
		OriginalImage:  userData.OriginalImage,
		Token:          token,
	}
	return tk
}
