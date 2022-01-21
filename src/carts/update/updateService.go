package carts

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")

func UpdateService(cart entity.AddToCart) string {
	// query filters
	filter := bson.M{"user_id": cart.UserId, "products.product_id": cart.ProductId}
	update := bson.M{"products.$[item].quantity": cart.Quantity}
	arrayFilter := bson.M{"item.product_id": cart.ProductId}

	// query db
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

	// check error
	if res.Err() != nil {
		panic(res.Err())
	}

	// return if no error
	return "Success"
}
