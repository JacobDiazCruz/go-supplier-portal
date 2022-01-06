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

	// get user by email
	userData := getUser.GetEmail(claims.Username)

	// return response
	tk := &TokenIdentity{
		Username:       userData.FirstName + userData.LastName,
		Email:          userData.Email,
		ThumbnailImage: userData.ThumbnailImage,
		OriginalImage:  userData.OriginalImage,
		Token:          token,
	}
	return tk
}
