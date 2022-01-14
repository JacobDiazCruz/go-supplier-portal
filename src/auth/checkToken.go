package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	getUser "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var myCollection *mongo.Collection = database.OpenCollection(database.Client, "blacklists")

func CheckToken(ctx *gin.Context) {
	if len(ctx.Request.Header["Authorization"]) <= 0 {
		fmt.Println("here error token")
	}
	token := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]

	fmt.Println(token)
	fmt.Println("here token")
	tokenStr := token
	claims := &Claims{}
	identity := &TokenIdentity{}

	// validate with claims
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Bad Request"})
		return
	}

	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
		return
	}

	// get user by email
	fmt.Println(claims.Email)
	fmt.Println("42424242424242")
	userData := getUser.GetEmail(claims.Email)

	fmt.Println(userData)
	fmt.Println("ewewewew")

	// check jwt if it is blacklisted
	query := bson.M{"token": token}
	err2 := myCollection.FindOne(context.TODO(), query).Decode(&identity)

	// return response
	tk := &TokenIdentity{
		Username:       userData.Username,
		Email:          userData.Email,
		ThumbnailImage: userData.ThumbnailImage,
		OriginalImage:  userData.OriginalImage,
		Token:          token,
	}

	if err2 != nil {
		// return if it is valid
		ctx.JSON(http.StatusOK, gin.H{"data": tk})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"data": "Invalid token"})
		return
	}
}
