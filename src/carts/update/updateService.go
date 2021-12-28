package carts

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")

func UpdateService(cart entity.ProductRequest, id string) string {
	// id to mongoId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	filter := bson.M{"_id": objID, "products.product_id": cart.ProductId}
	update := bson.M{"products.$[item].quantity": cart.Quantity}
	arrayFilter := bson.M{"item.product_id": cart.ProductId}
	// result, err := cartCollection.FindOneAndUpdate(
	// 	context.Background(),

	// )

	res := cartCollection.FindOneAndUpdate(context.Background(),
		filter,
		bson.M{"$set": update},
		options.FindOneAndUpdate().SetArrayFilters(
			options.ArrayFilters{
				Filters: []interface{}{
					arrayFilter,
				},
			},
		))

	if res.Err() != nil {
		panic("err")
		// log error
	}
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)

	// return
	return id
}
