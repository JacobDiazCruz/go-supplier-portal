package carts

import (
	"context"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	getProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/get"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")

func AddService(cart entity.AddToCart) string {
	// query products and add sales_information on the request
	// @TODO: Throw error if product id does not exist
	productDetails := getProduct.GetService(cart.ProductId, "")

	// @TODO: Query variants in

	// query
	result, err := cartCollection.UpdateOne(
		context.TODO(),
		bson.M{"user_id": cart.UserId},
		bson.M{
			"$push": bson.M{
				"products": bson.M{
					"product_id":        cart.ProductId,
					"variants":          cart.Variants,
					"name":              productDetails.Name,
					"thumbnail_image":   productDetails.ThumbnailImage,
					"original_image":    productDetails.OriginalImage,
					"sales_information": productDetails.SalesInformation,
					"quantity":          cart.Quantity,
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// return
	return cart.UserId.Hex()
}
