package carts

import (
	"context"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")

func ClearService(cart entity.AddToCart) string {
	// query db
	result, err := cartCollection.UpdateOne(
		context.TODO(),
		bson.M{"user_id": cart.UserId},
		bson.M{
			"$set": bson.M{
				"products": []entity.ProductResponse{},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// return if no error
	return cart.ProductId
}
