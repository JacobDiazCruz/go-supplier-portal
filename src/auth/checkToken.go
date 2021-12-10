package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var myCollection *mongo.Collection = database.OpenCollection(database.Client, "blacklists")

func CheckToken(ctx *gin.Context) {
	token := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	claims := &Claims{}
	identity := &TokenIdentity{}

	// if token exists
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
	// 	}
	// 	ctx.JSON(http.StatusOK, gin.H{"data": "Success"})
	// }

	// validate with claims
	tkn, err := jwt.ParseWithClaims(token, claims,
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

	// check jwt if it is blacklisted
	query := bson.M{"token": token}
	err2 := myCollection.FindOne(context.TODO(), query).Decode(&identity)
	if err2 != nil {
		// return if it is valid
		ctx.JSON(http.StatusOK, gin.H{"data": token})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"data": "Invalid token"})
		return
	}
}
