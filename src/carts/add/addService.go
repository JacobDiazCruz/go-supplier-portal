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

	// get selected variant details
	variantsRequest := []interface{}{}
	for _, productVariant := range productDetails.Variants {
		for _, cartVariant := range cart.Variants {
			if productVariant.ID == cartVariant.VariantId {
				for _, productVariantOption := range productVariant.Options {
					if cartVariant.VariantOptionId == productVariantOption.ID {
						variantEntity := entity.Variant{}
						variantEntity.ID = productVariant.ID
						variantEntity.Name = productVariant.Name
						variantEntity.Option = productVariantOption
						variantsRequest = append(variantsRequest, variantEntity)
					}
				}
			}
		}
	}

	// query
	result, err := cartCollection.UpdateOne(
		context.TODO(),
		bson.M{"user_id": cart.UserId},
		bson.M{
			"$push": bson.M{
				"products": bson.M{
					"_id":                   cart.ProductId,
					"variants":              variantsRequest,
					"slug":                  productDetails.Slug,
					"name":                  productDetails.Name,
					"files":                 productDetails.Files,
					"thumbnail_display_url": productDetails.ThumbnailDisplayUrl,
					"sales_information":     productDetails.SalesInformation,
					"quantity":              cart.Quantity,
					"audit_log":             productDetails.AuditLog,
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
