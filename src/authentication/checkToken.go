package authentication

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var myCollection *mongo.Collection = database.OpenCollection(database.Client, "blacklists")

func CheckToken(ctx *gin.Context) {
	cookie, err := ctx.Cookie("token")

	// if token exists
	if err != nil {
		if err == http.ErrNoCookie {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		}
		ctx.JSON(http.StatusOK, gin.H{"data": "Success"})
	}

	tokenStr := cookie
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

	// check jwt if it is blacklisted
	query := bson.M{"token": tokenStr}
	err2 := myCollection.FindOne(context.TODO(), query).Decode(&identity)
	if err2 != nil {
		// return if it is valid
		ctx.JSON(http.StatusOK, gin.H{"data": tokenStr})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"data": "Invalid token"})
		return
	}

}
