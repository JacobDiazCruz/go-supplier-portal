package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/authentication"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func LoginController(ctx *gin.Context) {
	fmt.Println("Im here")
	var credentials entity.Credentials

	err := ctx.BindJSON(&credentials)

	// if payload is invalid
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Bad Request"})
		return
	}

	expectedPassword, ok := users[credentials.Username]
	// if credentials is incorrect
	if !ok || expectedPassword != credentials.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		return
	}

	tk := auth.TokenIdentity{credentials.Username}
	signToken := auth.SignToken(tk)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString(jwtKey)

	// // if jwt signature is invalid
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error"})
	// 	return
	// }

	// ctx.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Sucess Login", "data": signToken})
}
