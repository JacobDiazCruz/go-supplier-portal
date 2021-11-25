package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Bad Req"})
		return
	}

	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauth"})
		return
	}
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &entity.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Println(tokenString)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error"})
		return
	}

	ctx.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Sucess Login", "data": tokenString})
}
