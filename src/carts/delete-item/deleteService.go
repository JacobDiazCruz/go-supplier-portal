package carts

import (
	"context"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")

func DeleteService(cart entity.ProductRequest, id string) string {
	// convert string id to mongoId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query db
	result, err := cartCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$pull": bson.M{
				"products": bson.M{"product_id": cart.ProductId},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("update service here id ^")

	// return if no error
	return id
}
