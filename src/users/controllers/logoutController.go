package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

var myCollection *mongo.Collection = database.OpenCollection(database.Client, "blacklists")

func LogoutController(ctx *gin.Context) {
	// check token and return
	ct, err := auth.GetToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
	}
	jwt := entity.TokenIdentity{ct}

	// blacklist to db
	result, err2 := myCollection.InsertOne(context.TODO(), jwt)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result)
	ctx.JSON(http.StatusOK, gin.H{"data": jwt})
}
